package goartm

import "sort"

//import "fmt"

type WeightedObject struct {
	Weight float32
	Object string
	ID     int
}

type WeightedObjects []WeightedObject

func (s WeightedObjects) Len() int           { return len(s) }
func (s WeightedObjects) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s WeightedObjects) Less(i, j int) bool { return s[i].Weight > s[j].Weight }

func GetTopTopicTokens(topicID int, tm *TopicModel, n int) (WeightedObjects, int) {
	var topicWords WeightedObjects

	for word_id, topics := range tm.TopicIndices {
		for i, topic := range topics.Value {
			if topicID == int(topic) {
				topicWords = append(topicWords, WeightedObject{ID: int(word_id), Weight: tm.TokenWeights[word_id].Value[i]})
				break
			}
		}
	}
	var ret WeightedObjects
	if len(topicWords) > 0 {
		sort.Sort(topicWords)
		if len(topicWords) < n {
			n = len(topicWords)
		}
		ret = make(WeightedObjects, n)

		for i, tw := range topicWords[:n] {
			ret[i] = tw
			ret[i].Object = tm.GetToken()[tw.ID]
		}
		return ret, len(topicWords)
	}
	return WeightedObjects{}, 0
}

func GetTopicTokens(topicID int, tm *TopicModel) map[string]float32 {
	topicWords := make(map[string]float32)
	for word_id, topics := range tm.TopicIndices {
		for i, topic := range topics.Value {
			if topicID == int(topic) {
				topicWords[tm.Token[word_id]] = tm.TokenWeights[word_id].Value[i]
				break
			}
		}
	}
	return topicWords
}

//NewBatchFromData
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
		title := ids[i] + " " + title[i]
		item.Title = &title
		// generate BOW representation
		for token, tf := range localDic {
			item.TokenId = append(item.TokenId, dic[token])
			item.TokenWeight = append(item.TokenWeight, tf)
		}
		batch.Item = append(batch.Item, item)
	}
	return batch
}
