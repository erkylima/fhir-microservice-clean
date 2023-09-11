package provider

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/go-playground/validator/v10"

	"github.com/qbem-repos/patient-service/internal/patient/provider/dbmongo"
	"github.com/qbem-repos/patient-service/pkg/config"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PatientProvider struct {
	collection *mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

type PatientChangeData struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type PatientPullFilter struct {
	Offset int64
	Limit  int64
}

func NewPatientProvider() *PatientProvider {
	ctx, cancel := context.WithCancel(context.Background())
	mongo_uri := config.MongoUri()
	uri, _ := url.Parse(mongo_uri)
	collection, err := dbmongo.NewCollection(uri.String(), "patients", "patients_collection")
	if err != nil {
		defer cancel()
		log.Printf("Error: %v", err)
		return nil
	}
	return &PatientProvider{
		collection: collection,
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (p *PatientProvider) Push(m *fhir.Patient) (string, error) {
	defer p.collection.Database().Client().Disconnect(p.ctx)

	paciente := fhir.Patient{}
	var opts = options.FindOne()
	var filter = bson.D{{Key: "registryCode", Value: m.Id}}

	errDecode := p.collection.FindOne(context.TODO(), filter, opts).Decode(&paciente)
	if errDecode == nil {
		return "", fmt.Errorf("register code %s is already registered", *m.Id)
	}

	var val = validator.New()
	err := val.Struct(m)

	if err != nil {
		return "", err.(validator.ValidationErrors)
	}

	_, err = p.collection.InsertOne(context.TODO(), m)

	if err != nil {
		return "", fmt.Errorf("register code %s is already registered", *m.Id)
	}

	return *m.Id, nil
}

func (p *PatientProvider) Count(filter *PatientPullFilter) int64 {
	data := bson.D{{}}
	var opts = new(options.CountOptions)

	if filter.Limit != 0 {
		opts = opts.SetLimit(filter.Limit)
	}

	if filter.Offset != 0 {
		opts = opts.SetSkip(filter.Offset)
	}
	total, _ := p.collection.CountDocuments(context.TODO(), data, opts)
	return total
}

func (p *PatientProvider) UpdateOne(slug string, key string, value any) error {
	var err error
	var opts = options.Update().SetUpsert(true)
	var filter = bson.D{{Key: "id", Value: slug}}
	var update = bson.D{{Key: "$set", Value: bson.D{{Key: key, Value: value}}}}
	result, err := p.collection.UpdateOne(p.ctx, filter, update, opts)

	if err != nil {
		log.Println(result)
		return errors.New("não foi possível atualizar o paciente")
	}

	return nil
}

// PullOne realiza consulta de paciente por id
func (p *PatientProvider) PullOne(id string) (*fhir.Patient, error) {
	defer p.collection.Database().Client().Disconnect(p.ctx)

	var patient *fhir.Patient
	var opts = options.FindOne()
	var filter = bson.D{{Key: "_id", Value: id}}
	err := p.collection.FindOne(p.ctx, filter, opts).Decode(&patient)
	if err != nil {
		return nil, err
	}
	log.Println("GetPatient successfully")
	return patient, nil
}

// Pull realiza consulta de pacientes
func (p *PatientProvider) Pull(filter PatientPullFilter) (*[]fhir.Patient, error) {
	var err error
	var patients = []fhir.Patient{}
	var opts = options.Find()
	var doc = bson.D{{}}

	if filter.Limit != 0 {
		opts = opts.SetLimit(filter.Limit)
	}

	if filter.Offset != 0 {
		opts = opts.SetSkip(filter.Offset)
	}

	//opts = opts.SetSort(bson.D{{Key: "fullname", Value: 1}, {Key: "id", Value: 1}})

	cur, err := p.collection.Find(p.ctx, doc, opts)

	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("not is possible find patients")
	}

	patients, err = p.iterator(cur, err, patients)
	defer p.collection.Database().Client().Disconnect(p.ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &patients, nil
}

// interator
func (*PatientProvider) iterator(cur *mongo.Cursor, err error, patients []fhir.Patient) ([]fhir.Patient, error) {
	for cur.Next(context.TODO()) {
		var patient = new(fhir.Patient)
		if err = cur.Decode(patient); err != nil {
			log.Fatalln(err.Error())
			continue
		}

		patients = append(patients, *patient)
	}
	return patients, err
}

func (p *PatientProvider) ListAllPatients() ([]*fhir.Patient, error) {
	var patients []*fhir.Patient
	cursor, err := p.collection.Find(p.ctx, bson.D{{}})
	if err != nil {
		log.Printf("Error on ListAllPatients: %v", err)
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var patient *fhir.Patient
		err = cursor.Decode(&patient)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	log.Println("ListAllPatients successfully")
	return patients, nil
}

func (p *PatientProvider) DeleteOneBySlug(slug string) (int, error) {
	var filter = bson.D{{Key: "slug", Value: slug}}
	defer p.collection.Database().Client().Disconnect(p.ctx)

	result, err := p.collection.DeleteOne(p.ctx, filter)

	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("nao foi possível deletar o paciente")
	}
	if result.DeletedCount == 0 {
		return 0, errors.New("paciente nao encontrado")
	}

	return int(result.DeletedCount), nil
}
