# Parse & Go

This is a common utility for the use of parse action of configs of your application. 

While we develop our applications, we need to have changeable values for our application to ease the management. For this purpose, we use configurations.

These configurations may be defined on your application directory path, or in a remote config server. 

These configurations also can be in many formats like _JSON, YAML, PROPERTIES, ENV_ etc.

This module enables you to read your configurations in your local or in your remote config server.

I will be giving the details about how to use this utility to read your configurations in your applications.

## Steps

### Step 1

`go get github.com/yukselcodingwithyou/parseandgo@latest`

First, you need to get module to your application. 

### Step 2

`import configparser "github.com/yukselcodingwithyou/parseandgo"`

Import the downloaded utility to your go files, as visualized above.

### Step 3

Now, you can create your parser with the types defined below.

 - **_JSON_**
 - **_YML_**
 - **_PROPERTIES_**
 - **_ENV_**

Also, you need to give an address for your file, can be a remote file in a server, for example (Amazon S3), or in your local directory.

## Examples 

I will give examples of every format type with url and local file.

1. JSON EXAMPLE

- Local File


    {
        "fileFormat": "json",
        "name": "yuksel",
        "surname": "ozdemir",
        "otherInfo": {
            "age": 26,
            "city": "Ankara",
            "student": false
        }
    }


- Code Example


       jsonParser := configparser.NewParser(configparser.JSON, "example.json")
       jsonConfig := configparser.Parse(jsonParser)
        
       student, err := jsonConfig.Value("otherInfo", "student").Bool()
       if err == nil {
         fmt.Println(*student)
       }

       age, err := jsonConfig.Value("otherInfo", "age").Int()
       if err == nil {
         fmt.Println(*age)
       }

       city, err := jsonConfig.Value("otherInfo", "city").String()
       if err == nil {
         fmt.Println(*city)
       }

       surname, err := jsonConfig.Value("surname").String()
       if err == nil {
         fmt.Println(*surname)
       }

       name, err := jsonConfig.Value("name").String()
       if err == nil {
         fmt.Println(*name)
       }

       fileFormat, err := jsonConfig.Value("fileFormat").String()
       if err == nil {
         fmt.Println(*fileFormat)
       }

2. YML EXAMPLE

- URL: `https://configserverexample.s3.eu-central-1.amazonaws.com/example.yml`

- Code Example


    yamlParser := configparser.NewParser(configparser.YAML, "https://configserverexample.s3.eu-central-1.amazonaws.com/example.yml")
    yamlConfig := configparser.Parse(yamlParser)

    student, err := yamlConfig.Value("other", "info", "person", "student").Bool()
	if err == nil {
		fmt.Println(*student)
	}
	age, err := yamlConfig.Value("other", "info", "person", "age").Int()
	if err == nil {
		fmt.Println(*age)
	}
	city, err := yamlConfig.Value("other", "info", "person", "city").String()
	if err == nil {
		fmt.Println(*city)
	}

	surname, err := yamlConfig.Value("person", "surname").String()
	if err == nil {
		fmt.Println(*surname)
	}

	name, err := yamlConfig.Value("person", "name").String()
	if err == nil {
		fmt.Println(*name)
	}

	fileFormat, err := yamlConfig.Value("file", "format").String()
	if err == nil {
		fmt.Println(*fileFormat)
	}
  
3. PROPERTIES EXAMPLE

- Local File


    file.format=properties
    person.name=yuksel
    person.surname=ozdemir
    other.info.person.age=26
    other.info.person.city=Ankara
    other.info.person.student=false
    

- Code Example


        propertiesParser := configparser.NewParser(configparser.PROPERTIES, "example.properties")
        propertiesConfig := configparser.Parse(propertiesParser)

        student, err := propertiesConfig.Value("other.info.person.student").Bool()
        if err == nil {
            fmt.Println(*student)
        }
        age, err := propertiesConfig.Value("other.info.person.age").Int()
        if err == nil {
            fmt.Println(*age)
        }
        city, err := propertiesConfig.Value("other.info.person.city").String()
        if err == nil {
            fmt.Println(*city)
        }
        
        surname, err := propertiesConfig.Value("person.surname").String()
        if err == nil {
            fmt.Println(*surname)
        }
        
        name, err := propertiesConfig.Value("person.name").String()
        if err == nil {
            fmt.Println(*name)
        }
        
        fileFormat, err := propertiesConfig.Value("file.format").String()
        if err == nil {
            fmt.Println(*fileFormat)
        }
        
4. ENV EXAMPLE

- URL: `https://configserverexample.s3.eu-central-1.amazonaws.com/example.env`

- Code Example


    envParser := configparser.NewParser(configparser.ENV, "https://configserverexample.s3.eu-central-1.amazonaws.com/example.env")
    envConfig := configparser.Parse(envParser)

    student, err := envConfig.Value("student").Bool()
	if err == nil {
		fmt.Println(*student)
	}
	age, err := envConfig.Value("age").Int()
	if err == nil {
		fmt.Println(*age)
	}
	city, err := envConfig.Value("city").String()
	if err == nil {
		fmt.Println(*city)
	}

	surname, err := envConfig.Value("surname").String()
	if err == nil {
		fmt.Println(*surname)
	}

	name, err := config.envConfig("name").String()
	if err == nil {
		fmt.Println(*name)
	}

	fileFormat, err := envConfig.Value("fileFormat").String()
	if err == nil {
		fmt.Println(*fileFormat)
	}


Supported types for a value in config are listed below:

- **_String_**
- **_Float_**
- **_Int_**
- **_Bool_**


## Conclusion

You can use this config parser, to parse any of your files with supported file formats. 

For any type of questions, please open issues :)

Happy Coding :)