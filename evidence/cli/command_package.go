package cli

import (
	"github.com/jfrog/jfrog-cli-artifactory/evidence"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	coreConfig "github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
)

type evidencePackageCommand struct {
	ctx     *components.Context
	execute execCommandFunc
}

func NewEvidencePackageCommand(ctx *components.Context, execute execCommandFunc) EvidenceCommands {
	return &evidencePackageCommand{
		ctx:     ctx,
		execute: execute,
	}
}

func (epc *evidencePackageCommand) CreateEvidence(ctx *components.Context, serverDetails *coreConfig.ServerDetails) error {
	err := epc.validateEvidencePackageContext(ctx)
	if err != nil {
		return err
	}

	createCmd := evidence.NewCreateEvidencePackage(
		serverDetails,
		epc.ctx.GetStringFlagValue(predicate),
		epc.ctx.GetStringFlagValue(predicateType),
		epc.ctx.GetStringFlagValue(markdown),
		epc.ctx.GetStringFlagValue(key),
		epc.ctx.GetStringFlagValue(keyAlias),
		epc.ctx.GetStringFlagValue(packageName),
		epc.ctx.GetStringFlagValue(packageVersion),
		epc.ctx.GetStringFlagValue(packageRepoName))
	return epc.execute(createCmd)
}

func (epc *evidencePackageCommand) validateEvidencePackageContext(ctx *components.Context) error {
	if !ctx.IsFlagSet(packageVersion) || assertValueProvided(ctx, packageVersion) != nil {
		return errorutils.CheckErrorf("--%s is a mandatory field for creating a Package evidence", packageVersion)
	}
	if !ctx.IsFlagSet(packageRepoName) || assertValueProvided(ctx, packageRepoName) != nil {
		return errorutils.CheckErrorf("--%s is a mandatory field for creating a Package evidence", packageRepoName)
	}
	return nil
}
