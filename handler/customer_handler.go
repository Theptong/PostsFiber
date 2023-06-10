package handler

import (
	"workshop/service"
)

type customerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{customerService: custSrv}
}


// func (h customerHandler) GETABC(c *fiber.Ctx) error {
// 	app := fiber.New()
// 	todoForm := struct {
// 		Text string `json:"text"`
// 	}{}
// 	app.Post("/", func(c *fiber.Ctx) error {

// 		c.BodyParser(&todoForm) // "{"Text":"do something"}"
// 		return c.JSON(todoForm)
// 	})

// 	return c.JSON(todoForm)
// }

// func (h customerHandler) GetCustomers(c *fiber.Ctx) error {
// 	customer, err := h.customerService.GetCustomerService()
// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 		fmt.Fprint(c, err)
// 		// return
// 	}
// 	// c.BodyParser(&customer)
// 	//การสร้าง json ส่งออกไปยัง หน้าจอ ให้แสดงผล
// 	fmt.Println("c.JSON(&customer)::", c.JSON(&customer))
// 	return c.JSON(&customer)
// }

// func (h customerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
// 	customerID, _ := strconv.Atoi(mux.Vars(r)["customer_id"])
// 	fmt.Println("customerID::", customerID)
// 	customer, err := h.customerService.GetCustomerServiceById(customerID)

//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			fmt.Fprint(w, err)
//			return
//		}
//		//การสร้าง json ส่งออกไปยัง หน้าจอ ให้แสดงผล
//		w.Header().Set("content-type", "application/json")
//		json.NewEncoder(w).Encode(customer)
//		// fmt.Println(customer)
//	}
//
