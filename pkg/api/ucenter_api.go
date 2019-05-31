package api

// func (r *XRouter) strategies(c *gin.Context) {
// 	c.JSON(200, nil)
// }

// func (r *XRouter) userList(c *gin.Context) {
// 	c.JSON(200, nil)
// }

// func (r *XRouter) user(c *gin.Context) {
// 	str := `{"success":true,"user":{"id":0,"username":"admin","permissions":{"role":"admin"},"avatar":"https://pbs.twimg.com/profile_images/835224725815246848/jdMBCxHS.jpg"}}`
// 	header := c.Writer.Header()
// 	header["Content-Type"] = []string{"application/json; charset=utf-8"}
// 	c.Writer.Write([]byte(str))
// 	c.Writer.WriteHeader(200)
// }

// func (r *XRouter) routes(c *gin.Context) {
// 	data, err := ioutil.ReadFile("routes.json")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	header := c.Writer.Header()
// 	header["Content-Type"] = []string{"application/json; charset=utf-8"}
// 	c.Writer.Write(data)
// 	c.Writer.WriteHeader(200)
// }
