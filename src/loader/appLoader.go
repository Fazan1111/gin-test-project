package loader

import "learnGin/src/loader/mongo"

func AppLoader() {
	mongo.LoadMongoDB()
}
