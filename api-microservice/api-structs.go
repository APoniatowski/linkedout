package main

type WelcomeJSON struct {
	Message	string `json:"message"`
}

type ContactJSON struct {
	Email		string `json:"email"`
	LinkedinURL string `json:"linkedinurl"`
	Phone		string `json:"phone"`
}