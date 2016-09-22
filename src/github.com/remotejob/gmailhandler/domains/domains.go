package domains

type Email struct {
	To      string
	Subject string
	Body    string
}

type Client struct {
	Phone string
	Email string
	Skype string

}


type ServerConfig struct {
	Login struct {
		Glogin string
		
	}
	Pass struct {
		Gpass string
		
	}


}
