package transform

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/pkg"
)

type ResultConfigFn func(*Transform, interface{}, *models.Card) interface{}

type Transform struct {
	table      models.TrelloTable
	labelsMap  map[string]*models.Label
	briefFnMap map[string]ResultConfigFn
	result     map[string]interface{}
}

func New(rawData []byte) *Transform {
	var t Transform
	t.initTransform()

	err := trello.UnmarshalJson(rawData, &t.table)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	t.allocateLabelsMap()
	return &t
}

func (t *Transform) initTransform() {
	t.labelsMap = make(map[string]*models.Label)
	t.briefFnMap = make(map[string]ResultConfigFn)
	t.result = make(map[string]interface{})
}

func (t *Transform) allocateLabelsMap() {
	for i, v := range t.table.Labels {
		t.labelsMap[v.ID] = &t.table.Labels[i]
	}
}

func (t *Transform) ResultConfig(key string, fn ResultConfigFn) {
	t.briefFnMap[key] = fn
}

func (t *Transform) TransformFromTrello() {
	for _, card := range t.table.Cards {
		for key, fn := range t.briefFnMap {
			t.result[key] = fn(t, t.result[key], &card)
		}
	}
}

func (t *Transform) GetAllResult() map[string]interface{} {
	return t.result
}

func (t *Transform) GetResult(key string) interface{} {
	return t.result[key]
}

func ExtractReferenceFn(tr *Transform, preValue interface{}, c *models.Card) interface{} {
	arr, ok := preValue.([]string)
	if !ok {
		arr = []string{}
	}
	reg, err := regexp.Compile(`[[][\S\s]+[]][(][\S]+[)]`)
	if err != nil {
		log.Printf("Extract reference error: %s \n", err.Error())
		return nil
	}
	targets := reg.FindAllString(c.Desc, -1)
	arr = append(arr, targets...)

	return arr
}

func CountLabelsFn(tr *Transform, preValue interface{}, c *models.Card) interface{} {
	labelsMap, ok := preValue.(map[string]int)
	if !ok {
		labelsMap = make(map[string]int)
	}
	for _, id := range c.IDLabels {
		v, ok := tr.labelsMap[id]
		if ok {
			labelsMap[v.Name]++
		}
	}

	return labelsMap
}

func CardBriefFn(tr *Transform, preValue interface{}, c *models.Card) interface{} {
	if c.Due == "" || c.Name == "" {
		return nil
	}

	v, ok := preValue.([]string)

	if !ok {
		v = []string{}
	}

	t, err := time.Parse(time.RFC3339, c.Due)
	if err != nil {
		log.Printf("[Transform] Parse time error %s", err.Error())
		return nil
	}

	item := fmt.Sprintf("%d %s %d  %s", t.Year(), t.Month().String(), t.Day(), c.Name)

	v = append(v, item)
	return v
}
