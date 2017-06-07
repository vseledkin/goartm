package goartm

import "sort"

type WeightedObject struct {
	Weight float32
	Object string
	ID     int
}

type WeightedObjects []WeightedObject

func (s WeightedObjects) Len() int           { return len(s) }
func (s WeightedObjects) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s WeightedObjects) Less(i, j int) bool { return s[i].Weight > s[j].Weight }

func GetTopTopicTokens(topicID int, tokens []string, tetaMatrix []*FloatArray, n int) (WeightedObjects, int) {
	var topicWords WeightedObjects
	for i, tokenWeight := range tetaMatrix {
		value := tokenWeight.GetValue()[topicID]
		if value > 0 {
			wo := WeightedObject{ID: i, Weight: value}
			topicWords = append(topicWords, wo)
		}
	}

	sort.Sort(topicWords)

	ret := make(WeightedObjects, n)
	for i, tw := range topicWords[:n] {
		ret[i] = tw
		ret[i].Object = tokens[tw.ID]
	}
	return ret, len(topicWords)
}

func NewBatchFromData(ids []string, title []string, data [][]string) *Batch {
	// batch.id must be set to a unique GUID in a format of 00000000-0000-0000-0000-000000000000.
	dic := make(map[string]int32)
	for _, doc := range data {
		for _, token := range doc {
			_, ok := dic[token]
			if !ok {
				dic[token] = int32(len(dic))
			}
		}
	}
	batch := NewBatch()
	// fill batch with tokens
	batch.Token = make([]string, len(dic))
	for token := range dic {
		batch.Token[dic[token]] = token
	}
	// fill documents
	for i, doc := range data {
		localDic := make(map[string]float32)
		for _, token := range doc {
			localDic[token]++
		}
		id := int32(i)
		item := NewItem()
		item.Id = &id // document id
		title := title[i]
		item.Title = &title
		id_str := ids[i]
		fieldName := "ID"
		item.Field = []*Field{&Field{Name: &fieldName, StringValue: &id_str}}
		// generate BOW representation

		for token, tf := range localDic {
			item.TokenId = append(item.TokenId, dic[token])
			item.TokenWeight = append(item.TokenWeight, tf)
		}
		batch.Item = append(batch.Item, item)
	}
	return batch
}
