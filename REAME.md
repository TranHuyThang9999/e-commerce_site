
func (t *ControllerUser) AddProfile(c *gin.Context) {
	var req models.ImagesReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	files, err := getUploadedFiles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Files = files

	resp, err := t.ctl.AddProfile(c, &req)
	if err != nil {
		c.JSON(200, resp)
		return
	}
	c.JSON(200, resp)
}