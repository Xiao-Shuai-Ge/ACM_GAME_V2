package jhinih_retriever

import (
	"ACM_GAME_V2/common/ais/ai_utils/RAG/client"
	"context"
	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino-ext/components/retriever/milvus"
)

func NewArkRetriever(ctx context.Context, embedder *ark.Embedder) *milvus.Retriever {
	client.InitClient()
	retriever, err := milvus.NewRetriever(ctx, &milvus.RetrieverConfig{
		Client:      client.MilvusCli,
		Collection:  "test",
		Partition:   nil,
		VectorField: "vector",
		OutputFields: []string{
			"id",
			"content",
			"metadata",
		},
		TopK:      4,
		Embedding: embedder,
	})
	if err != nil {
		panic(err)
	}
	return retriever
}
