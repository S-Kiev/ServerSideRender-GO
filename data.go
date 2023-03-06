package main

type clases struct {
	Titulo   string
	Duracion string
}

type modulo struct {
	Titulo      string
	Descripcion string
	Clases      []clases
}

type cursos struct {
	Slug        string
	Img         string
	Titulo      string
	Nombre      string
	Descripcion string
	Promedio    float64
	Profesor    string
	ProfesorImg string
	Precio      float64
	Modulos     []modulo
}

type gridPage struct {
	InternalTemplate string
	Cursos           []cursos
}

type cursosPage struct {
	InternalTemplate string
	Cursos           cursos
}

func loadGrilla() []cursos {
	return []cursos{
		{
			Slug:        "go",
			Img:         "https://edteam-media.s3.amazonaws.com/cursos/big/91e149d0-961a-4594-a8ff-0a625be9cdd2.png",
			Titulo:      "Go desde cero",
			Nombre:      "Go desde cero",
			Descripcion: "Adquiere los conocimientos básicos para iniciar a programar con Go.",
			Promedio:    4.7,
			Profesor:    "Alejo Rodriguez",
			ProfesorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/952327c3-2bd9-41d1-819e-9b5d7eb84c13.jpg",
			Precio:      30,
			Modulos: []modulo{
				{
					Titulo:      "Qué es Go",
					Descripcion: "Aprenderemos qué es el lenguaje Go",
					Clases: []clases{
						{
							Titulo:   "Historia",
							Duracion: "05:33",
						},
						{
							Titulo:   "Creadores",
							Duracion: "03:02",
						},
					},
				},
				{
					Titulo:      "Sintaxis",
					Descripcion: "Bases del lenguaje",
					Clases: []clases{
						{
							Titulo:   "Declaración de Variables",
							Duracion: "05:33",
						},
						{
							Titulo:   "Constantes",
							Duracion: "03:02",
						},
					},
				},
			},
		},
		{
			Slug:        "go-poo",
			Img:         "https://edteam-media.s3.amazonaws.com/cursos/big/8f37ebcc-16a1-4001-93e1-5f21893cd3cc.jpg",
			Titulo:      "POO con Go",
			Nombre:      "POO con Go",
			Descripcion: "Aprende a usar la programación orientada a objetos en Go",
			Promedio:    4.8,
			Profesor:    "Alejo Rodriguez",
			ProfesorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/952327c3-2bd9-41d1-819e-9b5d7eb84c13.jpg",
			Precio:      30,
		},
		{
			Slug:        "go-database",
			Img:         "https://edteam-media.s3.amazonaws.com/cursos/big/4d60ef81-2e58-457f-97c7-ee8847663985.jpg",
			Titulo:      "Bases de datos con Go",
			Nombre:      "Bases de datos con Go",
			Descripcion: "Aprende a integrar y usar las bases de datos en Go",
			Promedio:    4.8,
			Profesor:    "Alejo Rodriguez",
			ProfesorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/952327c3-2bd9-41d1-819e-9b5d7eb84c13.jpg",
			Precio:      30,
		},
		{
			Slug:        "go-testing",
			Img:         "https://edteam-media.s3.amazonaws.com/cursos/big/a9913502-8af2-400b-8095-7b78f52200dc.png",
			Titulo:      "Testing con Go",
			Nombre:      "Testing con Go",
			Descripcion: "Aprende a crear tests y medir la eficiencia de tus proyectos con Go",
			Promedio:    4.7,
			Profesor:    "Alexys Lozada",
			ProfesorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/66f015b2-0dfb-4ba9-bd0d-f7a7e1855275.jpeg",
			Precio:      24,
		},
	}
}

func getcursos(slug string) cursos {
	data := loadGrilla()
	for _, v := range data {
		if v.Slug == slug {
			return v
		}
	}

	return cursos{}
}
