package main

import (
	"log"
	"opengine/v2/internal/handler"
	"opengine/v2/internal/handler/router"
	middlerware "opengine/v2/internal/middleware"
	compRepo "opengine/v2/internal/repository/components"
	orgRepo "opengine/v2/internal/repository/organization"
	compServ "opengine/v2/internal/usecase/components"
	orgServ "opengine/v2/internal/usecase/organization"

	"opengine/v2/pkg/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// CARGA DE VARIABLES
	if err := godotenv.Load(); err != nil {
		log.Println("File config no found")
	}

	// carga de configuracion
	cfg := config.LoadConfig()
	dbConn := config.InitDB(cfg.DatabaseURL)
	defer dbConn.Close()

	// =========================================================================
	// INYECCIÓN DE DEPENDENCIAS (CAPA DE DATOS -> REPOSITORIOS)
	// =========================================================================
	authRepo := middlerware.NewAuthRepo(dbConn)
	user_repo := orgRepo.NewUserRepo(dbConn)

	// SESION DE COMPONENTES
	environments_repo := compRepo.NewEnvironmentRepo(dbConn)
	services_repo := compRepo.NewServicesRepo(dbConn)
	resources_repo := compRepo.NewRosourcesRepo(dbConn)

	// =========================================================================
	// INYECCIÓN DE DEPENDENCIAS (CAPA NEGOCIO -> SERVICIOS)
	// =========================================================================
	user_serv := orgServ.NewUserService(user_repo)

	// SESION DE COMPONENTES
	environments_serv := compServ.NewEnvironmentServ(environments_repo)
	services_serv := compServ.NewServicesServ(services_repo)
	resources_serv := compServ.NewRosourcesServ(resources_repo)

	// Servicio de Autenticación (Login/JWT)
	// Obtenemos la llave secreta desde el .env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "clave_por_defecto_solo_para_dev"
	}

	authServ := middlerware.NewAuthService(authRepo, jwtSecret)

	// =========================================================================
	// INYECCIÓN DE DEPENDENCIAS (CAPA DE PRESENTACION -> HANDLERS)
	// =========================================================================

	user_handler := handler.NewUserHandler(user_serv)

	environments_handler := handler.NewEnvironmentHandler(environments_serv)
	services_handler := handler.NewServicesHandler(services_serv)
	resources_handler := handler.NewRosourcesHandler(resources_serv)

	organizarionRouter := router.NewRouterOrganization(
		user_handler,
	)

	componentsRouter := router.NewRouterComponent(
		environments_handler,
		services_handler,
		resources_handler,
	)

	apiRouters := router.MainRouter{
		RouterOrganization: organizarionRouter,
		RouterComponent:    componentsRouter,
	}

	// =========================================================================
	// INICIALIZACIÓN DE GIN Y RUTAS CENTRALES
	// =========================================================================
	r := gin.Default()

	router.SetupRouter(r, authServ, apiRouters)

	// invocacion del servicio para los registros manuales
	router.RegisterUser(r, user_handler)

	// EJECUCION DEL SRV
	port := ":8084"
	log.Printf("Servidor en linea, port: %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Error al inicializar el srv: ", err)
	}
}
