package jhinih_indexer

import (
	"ACM_GAME_V2/common/ais/ai_utils/RAG/client"
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var collection = "test"

var fields = []*entity.Field{
	{
		Name:     "id",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "255",
		},
		PrimaryKey: true,
	},
	{
		Name:     "vector", // 确保字段名匹配
		DataType: entity.FieldTypeBinaryVector,
		TypeParams: map[string]string{
			"dim": "81920",
		},
	},
	{
		Name:     "content",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "8192",
		},
	},
	{
		Name:     "metadata",
		DataType: entity.FieldTypeJSON,
	},
}

func NewArkIndexer(ctx context.Context, embedder *ark.Embedder) *milvus.Indexer {
	client.InitClient()
	indexer, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     client.MilvusCli,
		Collection: collection,
		Fields:     fields,
		Embedding:  embedder,
		//DocumentConverter: floatDocumentConverter,//另外一种格式来存储于milvus的必备函数
	})
	if err != nil {
		fmt.Println("创建indexer失败: %v", err)
	}
	return indexer
}
