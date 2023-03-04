package chains

import (
	"context"
	omniflixnfttypes "github.com/OmniFlix/onft/types"
	"github.com/cosmos/cosmos-sdk/client"
)

type OmnfiFlixChain struct {
	ChainData
}

func (c OmnfiFlixChain) ListNFTs(ctx context.Context, clientCtx client.Context, query ListNFTsQuery) []NFT {
	nftQueryClient := omniflixnfttypes.NewQueryClient(clientCtx)

	request := &omniflixnfttypes.QueryOwnerONFTsRequest{
		DenomId: query.ClassReference,
		Owner:   query.Owner,
	}
	resp, err := nftQueryClient.OwnerONFTs(ctx, request)
	if err != nil {
		panic(err)
	}

	var nfts []NFT
	for _, collection := range resp.Collections {
		for _, nft := range collection.Onfts {
			nfts = append(nfts, NFT{
				ID:      nft.Id,
				ClassID: collection.GetDenom().Id,
			})
		}
	}

	return nfts
}

func (c OmnfiFlixChain) TransferNFT(ctx context.Context, clientCtx client.Context, fields TransferNFTFields) {
	panic("implement me")
}
