package Core

import (
	"Taxonim/Core/Types"
	"context"
	"net/http"
	"time"
)

var (
	InitEngineServices = func(this Types.Hammer) (*EngineServices, error) {

	}

	ReadTestData = func(TestDataConf map[string]Types.CsvConf) (map[string]Types.CsvData, error) {

	}

	CreateInitialCookies = func(Cookies []Types.CustomCookie) ([]*http.Cookie, error) {

	}
)

func (this *Engine) IsTestFailed() bool {

}

func (this *Engine) Init() (err error) {

}

func (this *Engine) Start() string {

}

func (this *Engine) RunWorkers(C int) {

}

func (this *Engine) RunWorker(ScenarioStartTime time.Time) {

}

func (this *Engine) RunAssertionsInEngine() bool {

}

func (this *Engine) Stop() {

}

func (this *Engine) GetMaxConcurrentIterCount() int {

}

func (this *Engine) InitReqCountArr() {

}

func (this *Engine) CreateManualReqCountArr() {

}

func (this *Engine) CreateLinearReqCountArr() {

}

func (this *Engine) CreateIncrementalReqCountArr() {

}

func (this *Engine) CreateWavedReqCountArr() {

}

func (this *Engine) CreateLinearDistArr() {

}

func (this *Engine) CreateIncrementalDistArr() {

}

func ArraySum(Steps []int) int {

}

func Reverse(S interface{}) {

}

func ParseRawCookie(Cookie string) []*http.Cookie {

}

func NewEngine(Ctx context.Context, H Types.Hammer, Service *EngineServices) (E Engine, err error) {

}
