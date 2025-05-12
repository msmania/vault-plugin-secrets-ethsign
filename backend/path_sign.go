package backend

import (
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathSign(b *backend) *framework.Path {
	return &framework.Path{
		Pattern:      "accounts/sign/message/" + framework.GenericNameRegex("name"),
		HelpSynopsis: "Sign a provided transaction object.",
		HelpDescription: `

    Sign a transaction object with properties conforming to the Ethereum JSON-RPC documentation.

    `,
		Fields: map[string]*framework.FieldSchema{
			"name": {Type: framework.TypeString},
			"message": {
				Type:        framework.TypeString,
				Description: "Message to sign in a hex string",
			},
			"hash": {
				Type:        framework.TypeString,
				Description: "Hash to sign in a hex string",
			},
		},
		ExistenceCheck: b.pathExistenceCheck,
		Callbacks: map[logical.Operation]framework.OperationFunc{
			// Cannot be read as it takes input
			logical.UpdateOperation: b.sign,
		},
	}
}
