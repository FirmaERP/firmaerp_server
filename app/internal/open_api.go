package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swaggest/swgui/v5emb"
)

func InitOpenAPI(ginGonic *GinGonic) {
	rootRoute := "/docs"
	fileRoute := "/files/docs"

	handler := v5emb.New(
		"FirmaERP",
		fmt.Sprintf("%s/openapi.yaml", fileRoute),
		rootRoute,
	)

	ginGonic.Engine.Static(fileRoute, "./dist/docs/")
	ginGonic.Engine.GET(fmt.Sprintf("%s/*any", rootRoute), gin.WrapH(handler))

	message := fmt.Sprintf(
		"[GIN-debug] Access the documentation at: http://127.0.0.1:%s%s",
		ginGonic.Port,
		rootRoute,
	)

	fmt.Println(message)
}
