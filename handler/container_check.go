package handler

import "web-IDE_manager/model"

func checkCreateContainer(req model.CreateContainerReqBody) bool {
	return !(req.UserID <= 0 || req.ImageID == "")
}
