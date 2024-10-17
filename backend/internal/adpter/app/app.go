package app

import "github.com/joho/godotenv"

func Build() {
	injectDependencies()
}

func injectDependencies() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
