To generate swagger files:
swagger generate server -q --server-package ../internal/pkg/restapi --main-package ../../cmd/novelgo-server --model-package ../internal/pkg/models --spec swagger.yml
(Typically swagger is installed here: ~/go/bin/swagger)
