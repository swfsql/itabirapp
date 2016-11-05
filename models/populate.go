package models

import (
	_ "fmt"
	_ "strconv"
	_ "strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func createPoster(u *User) {
	u.New()
}

func createData() {

	o := orm.NewOrm()

	o.Insert(&Institution_type{Name: "pesquisa"})
	o.Insert(&Institution_type{Name: "extensao"})
	o.Insert(&Institution_type{Name: "professor"})
	o.Insert(&Institution_type{Name: "estabelecimento"})
	o.Insert(&Institution_type{Name: "republica"})

	pesquisa, _ := GetInstitutionByName("pesquisa")
	extensao, _ := GetInstitutionByName("extensao")
	prof, _ := GetInstitutionByName("professor")
	estab, _ := GetInstitutionByName("estabelecimento")
	rep, _ := GetInstitutionByName("republica")

	o.Insert(&User{User_Type: "moderator", Name: "Felipe Cetrulo", NameTag: "moderador", NameIdTag: "moderador",
		Email: "fecetrulo@hotmail.com", Password: "123", Institution_type: nil, Institution_Description: "",
		Addr_Street: "", Addr_Number: "", Addr_Complement: "", Addr_Neighborhood: "", Addr_City: ""})

	o.Insert(&User{User_Type: "moderator", Name: "Thiago Machado", NameTag: "moderador", NameIdTag: "moderador2",
		Email: "swfsql@gmail.com", Password: "123", Institution_type: nil, Institution_Description: "",
		Addr_Street: "", Addr_Number: "", Addr_Complement: "", Addr_Neighborhood: "", Addr_City: ""})

	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Equipe de Robótica DRUMONSTERS",
		Email:                   "drumonsters@gmail.com",
		Password:                "rd",
		Institution_type:        &pesquisa,
		Institution_Tag:         "pesquisa",
		Institution_Description: "Criada em 2010 por alunos da UNIFEI - Itabira, a Drumonsters tem como objetivo o desenvolvimento tecnológico através da construção de robôs, participando de competições nacionais e futuramente internacionais.",
		Addr_Street:             "Rua Irmã Ivone Drummond",
		Addr_Number:             "200",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Distrito Industrial II",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "L.O.T.S.",
		Email:                   "gestao.lots@gmail.com",
		Password:                "lots",
		Institution_type:        &pesquisa,
		Institution_Tag:         "pesquisa",
		Institution_Description: "Equipe de competição de Aerodesign da Universidade Federal de Itajubá - Campus Itabira",
		Addr_Street:             "Avenida Ipiranga",
		Addr_Number:             "",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Major Lage de Baixo",
		Addr_City:               "Itabira",
	})

	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "UP Consultoria Jr.",
		Email:                   "contato@upconsultoriajr.com.br",
		Password:                "up",
		Institution_type:        &extensao,
		Institution_Tag:         "extensao",
		Institution_Description: "Empresa Jr. multidisciplinar da Universidade Federal de Itajubá - Campus Itabira.",
		Addr_Street:             "Rua Irmã Ivone Drummond",
		Addr_Number:             "200",
		Addr_Complement:         "1413",
		Addr_Neighborhood:       "Distrito Industrial II",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Ampari",
		Email:                   "ampariitabira@gmail.com",
		Password:                "amp",
		Institution_type:        &extensao,
		Institution_Tag:         "extensao",
		Institution_Description: "A Associação Municipal de Proteção Animal da Região de Itabira visa promover o Bem Estar animal e sua convivência harmoniosa na sociedade através da Guarda Responsável e a valorização do respeito à vida.",
		Addr_Street:             "Avenida Ipiranga",
		Addr_Number:             "",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Major Lage de Baixo",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Engenheiro sem fronteira",
		Email:                   "itabira@esf-brasil.org",
		Password:                "esf",
		Institution_type:        &extensao,
		Institution_Tag:         "extensao",
		Institution_Description: "Realizamos projetos de engenharia com finalidade social. Nossos pilares são: engenharia, educação, sustentabilidade e voluntariado.",
		Addr_Street:             "Rua Irmã Ivone Drummond",
		Addr_Number:             "200",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Distrito Industrial II",
		Addr_City:               "Itabira",
	})

	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Fabiana Costa Guedes",
		Email:                   "fabiana.costaguedes@unifei.edu.br",
		Password:                "fabi",
		Institution_type:        &prof,
		Institution_Tag:         "professor",
		Institution_Description: "Possui graduação em Ciência da Computação pela Pontifícia Universidade Católica de Minas Gerais (2000) e mestrado em Engenharia Elétrica pela Pontifícia Universidade Católica de Minas Gerais (2007). Atualmente é professor assistente da Universidade Federal de Itajubá - Campus Itabira. Tem experiência na área de Ciência da Computação, com ênfase em Engenharia de Software, atuando principalmente nos seguintes temas: qualidade, processo de software, desenvolvimento de sistemas, gerência de projetos e fábrica de software.",
		Addr_Street:             "Rua São Paulo",
		Addr_Number:             "377",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Bairro Amazonas",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Gustavo Henrique Oliveira Salgado",
		Email:                   "gusalg@unifei.edu.br",
		Password:                "salgado",
		Institution_type:        &prof,
		Institution_Tag:         "professor",
		Institution_Description: "Possui graduação em Matemática pela Universidade Federal de Minas Gerais (2003), mestrado em Matemática pela Universidade Federal de Minas Gerais (2006) e doutorado em Engenharia Elétrica pela Universidade Federal de Minas Gerais (2015). Atualmente é professor assistente da Universidade Federal de Itajubá. Tem experiência na área de Matemática, com ênfase em sistemas dinâmicos não lineares e sistemas de ordem não inteira.",
		Addr_Street:             "Rua Irmã Ivone Drumond",
		Addr_Number:             "200",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Distrito Industrial II",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Washington Batista Vieira",
		Email:                   "vieira@unifei.edu.br ",
		Password:                "violino",
		Institution_type:        &prof,
		Institution_Tag:         "professor",
		Institution_Description: "WASHINGTON BATISTA VIEIRA concluiu o doutorado em Engenharia Civil pela Universidade Federal de Viçosa (UFV) em 2015. Professor Adjunto na Universidade Federal de Itajubá (UNIFEI) - Campus Itabira. É engenheiro civil e mestre em engenharia civil pela UFV. Durante o curso de Doutorado, cursou disciplina no Programa de Pós-Graduação em Engenharia Civil da Universidade Federal de Ouro Preto (PROPEC/UFOP).",
		Addr_Street:             "Rua Irmã Ivone Drumond",
		Addr_Number:             "200",
		Addr_Complement:         "2441",
		Addr_Neighborhood:       "Distrito Industrial II",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Walter Aoiama Nagai",
		Email:                   "walternagai@unifei.edu.br",
		Password:                "nagai",
		Institution_type:        &prof,
		Institution_Tag:         "professor",
		Institution_Description: "Possui graduação em Bacharelado Em Ciência da Computação pela Universidade Federal de Mato Grosso do Sul (1997) e mestrado em Ciências da Computação e Matemática Computacional pela Universidade de São Paulo (2000). Atualmente é professor adjunto i, nível 1, ci da Universidade Federal de Itajubá. Tem experiência na área de Ciência da Computação, com ênfase em Sistemas de Informação, atuando principalmente nos seguintes temas: hipermídia, aprendizagem, gamificação, engenharia e ensino.",
		Addr_Street:             "Rua Irmã Ivone Drumond",
		Addr_Number:             "200",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Distrito Industrial II",
		Addr_City:               "Itabira",
	})

	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Cai&pira",
		Email:                   "rep_cai@gmail.com",
		Password:                "cai",
		Institution_type:        &rep,
		Institution_Tag:         "republica",
		Institution_Description: "Fundada em 2010 na cidade de São João del Rei, a República Cai&Pira teve seus dias de glória por aproximadamente 3 anos. Com a saída dos moradores, muitos acreditaram que este local para se reunir os amigos, conversar sobre quaisquer assuntos, tomar uma cerveja gelada nos fins de semana (na metade também, claro, porque não?) e sobretudo interagir, seria extinto. Mas o que não era mais muito esperado aconteceu, a Fênix nos deu o exemplo, e agora mais forte que nunca a República Cai&Pira está de volta! No coração de Belo Horizonte, mostrando à capital que aqui também é lugar de repúblicas estudantis e que nós estamos aqui para fazer a diferença.",
		Addr_Street:             "",
		Addr_Number:             "",
		Addr_Complement:         "",
		Addr_Neighborhood:       "",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Meteolate",
		Email:                   "rep_mete@gmail.com",
		Password:                "mete",
		Institution_type:        &rep,
		Institution_Tag:         "republica",
		Institution_Description: "Fundação em 2011",
		Addr_Street:             "Rua Honorina Machado",
		Addr_Number:             "208",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Amazonas",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "DNA",
		Email:                   "rep.dna@hotmail.com",
		Password:                "dna",
		Institution_type:        &rep,
		Institution_Tag:         "republica",
		Institution_Description: "República Dias e Noites Alcoolizados - UNIFEI - Itabira/MG. Fundação em 2008. Festas Open-Bar e Espetáculos.",
		Addr_Street:             "Rua Rio de Janeiro",
		Addr_Number:             "25",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Amazonas",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Amazonas",
		Email:                   "rep_ama@gmail.com",
		Password:                "ama",
		Institution_type:        &rep,
		Institution_Tag:         "republica",
		Institution_Description: "Casa ampla, próxima ao bretas, ponto de onibus e comércio. 8 quartos, garagem para 3 carros, internet 10mb. ",
		Addr_Street:             "Rua das Margaridas",
		Addr_Number:             "1113",
		Addr_Complement:         "",
		Addr_Neighborhood:       "São Pedro",
		Addr_City:               "Itabira",
	})

	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Bar do João",
		Email:                   "barjoao@gmail.com",
		Password:                "barjoao",
		Institution_type:        &estab,
		Institution_Tag:         "estabelecimento",
		Institution_Description: "Bar · Pub em Itabira",
		Addr_Street:             "Rua Água Santa",
		Addr_Number:             "191",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Centro",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Filé di gato",
		Email:                   "contato@filedigato.com.br",
		Password:                "file",
		Institution_type:        &estab,
		Institution_Tag:         "estabelecimento",
		Institution_Description: "A espeteria mais badalada de Itabira!",
		Addr_Street:             "Av Mauro Ribeiro Lage",
		Addr_Number:             "776",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Esplanada da Estação",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Bombar",
		Email:                   "bombar@gmail.com",
		Password:                "bombar",
		Institution_type:        &estab,
		Institution_Tag:         "estabelecimento",
		Institution_Description: "Itabira vai BOMBAR, Chegou o melhor Bar/choperia da região, cerveja e chopp´s gelados, os melhores petiscos e drinks.",
		Addr_Street:             "Av. Mauro Ribeiro",
		Addr_Number:             "675",
		Addr_Complement:         "Loja 01",
		Addr_Neighborhood:       "Vila São Joaquim",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:                    "Appricci",
		Email:                   "apricci@apricci.com.br",
		Password:                "apri",
		Institution_type:        &estab,
		Institution_Tag:         "estabelecimento",
		Institution_Description: "Pizzas de sabores especiais, assadas em forno a lenha e serviço a la cart a noite. Almoço self-service todos os dias. Ambiente climatizado e confortável.",
		Addr_Street:             "Avenida Mauro Ribeiro",
		Addr_Number:             "300",
		Addr_Complement:         "",
		Addr_Neighborhood:       "Major Lage de Baixo",
		Addr_City:               "Itabira",
	})
	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:             "Família Pires",
		Email:            "padaria@familiapires.com.br",
		Password:         "pires",
		Institution_type: &estab,
		Institution_Tag:  "estabelecimento",
		Institution_Description: `A Padaria Familia Pires traz para Itabira/MG o conceito das grandes padarias de São Paulo, aliando aos hábitos e costumes dos Itabiranos.
São mais de 1000m2 de área construída e mais de 60 vagas de estacionamento próprio.
O Horário de Funcionamento é de 2ª a 5ª feira, das 6h às 22h, e 6ª feira, sábado e domingo, das 6 às 24h.`,
		Addr_Street:       "Av Ver Osório Sampaio",
		Addr_Number:       "45",
		Addr_Complement:   "",
		Addr_Neighborhood: "Vila Santa Rosa",
		Addr_City:         "Itabira",
	})

	createPoster(&User{User_Type: "poster", IsAuthorized: true,
		Name:             "Mountain Baja",
		Email:            "mountainbaja@gmail.com",
		Password:         "mb",
		Institution_type: &pesquisa,
		Institution_Tag:  "pesquisa",
		Institution_Description: `Equipe de BAJA SAE da Universidade Federal de Itajubá - Campus Avançado de Itabira.

	O programa Baja SAE BRASIL é um desafio lançado aos estudantes de Engenharia que oferece a chance de aplicar na prática os conhecimentos adquiridos em sala de aula, visando incrementar sua preparação para o mercado de trabalho.

	Ao participar do programa Baja SAE, o aluno se envolve com um caso real de desenvolvimento de um veículo off road, desde sua concepção, projeto detalhado, construção e testes.`,
		Addr_Street:       "Rua Irmã Ivone Drumond",
		Addr_Number:       "200",
		Addr_Complement:   "",
		Addr_Neighborhood: "Distrito Industrial II",
		Addr_City:         "Itabira",
	})

	var posters []*User

	qs := o.QueryTable("user")
	_, _ = qs.Filter("User_Type", "poster").All(&posters)

	var user *User
	var post *Post

	// DRUMMONSTER
	user = posters[0]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Sem descrição",
		Text: `# Teste da Drummonsters
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[0]
	post = &Post{User: user,
		Title:       "Foto",
		Subtitle:    "360",
		Description: "Imagem da UNIFEI 360º",
		Text: `# Mais uma imagem da série UNIFEI 360º acaba de sair!
		Clique, arraste e descubra!
		#UNIFEI360 #Descubra #GoGoGoDrumonsters
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"foto", "360", "drone"})

	//L.O.T.S.
	user = posters[1]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da L.O.T.S.
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[1]
	post = &Post{User: user,
		Title:       "Projeto de Aeronaves",
		Subtitle:    "Introdução ao Projeto de Aeronaves",
		Description: "De quinta a sábado da última semana aconteceu na Unifei - Campus Itabira a 30ª edição do Curso de Introdução ao Projeto de Aeronaves, ministrado pelo Professor Luiz Eduardo Miranda.",
		Text: `De quinta a sábado da última semana aconteceu na Unifei - Campus Itabira a 30ª edição do Curso de Introdução ao Projeto de Aeronaves, ministrado pelo Professor Luiz Eduardo Miranda. A equipe LOTS Micro Aerodesign gostaria de agradecer a todos os patrocinadores que tornaram possível a realização desse evento:
			Pizzaria Romana, Supermercado Nova América, Família Pires, Gráfica Itabira, Distribuidora Jácome e Restaurante ComaBem.
			Muito obrigado pelo apoio!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"curso", "projeto", "aeronaves"})

	//UP Consultoria Jr
	user = posters[2]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da UP Consultoria Jr.
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[2]
	post = &Post{User: user,
		Title:       "Parceria",
		Subtitle:    "Parceria tigre",
		Description: "Nossa parceria com a Tigre representa muito para nossa comunidade!",
		Text: `Nossa parceria com a Tigre representa muito para nossa comunidade!
Estabelecendo conexões com grandes empresas, criamos um vínculo entre a universidade e o mercado de trabalho, onde trocaremos conhecimento, oportunidades e vivências.
A Expojob vem aí, e junto vem a Tigre para somar na transformação que iremos causar!
Expojob 2017, seja você o autor da sua história!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"parceria", "tigre", "expojob"})

	//Ampari
	user = posters[3]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Ampari
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[3]
	post = &Post{User: user,
		Title:       "Multirão",
		Subtitle:    "de castração",
		Description: "Nos dias 14 e 15 de Outubro foi realizado o II Mutirão de Castração na cidade! ",
		Text: `Mutirão de Castração
Nos dias 14 e 15 de Outubro foi realizado o II Mutirão de Castração na cidade! Sem o envolvimento da população, nada disso teria sido possível. É momento de agradecer...
As empresas: 
- Apricci, Restaurante do Gilson, Sabor Brasil e Oficina do pão: por viabilizar a alimentação da equipe de profissionais do castramóvel.
- Nossa Loja, Class, Crescer, Supermercado União, Fabiana Modas, Pedrinho Ouro e Prata, Império Uniformes e Diógenes: pelas doações em dinheiro para custeio do combustível, medicamentos, colares elizabetanos. 
- Agrocenter, Agroduarte e Casa do amigão pelos descontos em materiais e medicamentos.
Aos moto clubes: Impuros Brasil e Motoclubes Unidos de Itabira: pelas doações e pela disposição em ajudar!
Aos amigos, voluntários e parceiros da Associação: que se esforçaram para participar antes, durante e depois do evento!
Ao Deputado Noraldino pela presença durante o evento.
A ONG AJUDA pelo excelente trabalho realizado!
Estamos muito felizes com o retorno da população acreditando no trabalho que vem sendo desempenhado pela AMPARI. Pelas pessoas e pelos animais, ontem, hoje e sempre!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"multirao", "castração", "mobilização"})

	//Engenheiro sem fronteira
	user = posters[4]
	post = &Post{User: user,
		Title:       "Curso",
		Subtitle:    "Programa Cultura Empreendedora",
		Description: "Hoje alguns voluntários da ONG tiveram a oportunidade de participar do curso “Programa Cultura Empreendedora” do Sebrae.",
		Text: `Hoje alguns voluntários da ONG tiveram a oportunidade de participar do curso “Programa Cultura Empreendedora” do Sebrae. Aproveitamos cada momento e saímos com mais vontade de fazer acontecer! Gostaríamos de agradecer aos professores que nos receberam muito bem! Obrigada Lourdes, por ter nos passado todo conhecimento com tanta calma e clareza! E também gostaríamos de agradecer ao Sebrae pela oportunidade de desenvolver ainda mais os nossos projetos!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"curso", "programa", "sebrae"})

	user = posters[4]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Engenheiro sem fronteira
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	//Fabiana Costa Guedes
	user = posters[5]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Fabiana Costa Guedes
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[5]
	post = &Post{User: user,
		Title:       "Prova",
		Subtitle:    "de BD",
		Description: "A prova de BD do dia 25/04/2012 ocorrerá normalmente.",
		Text: `# A prova de BD do dia 25/04/2012 ocorrerá normalmente.!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"prova", "BD", "noticia"})

	//Gustavo Henrique Oliveira Salgado
	user = posters[6]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Gustavo Henrique Oliveira Salgado
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[6]
	post = &Post{User: user,
		Title:       "Prova",
		Subtitle:    "substituta",
		Description: "alteração prova substitura",
		Text: `Caros discentes,

embora havíamos pré-agendado a prova substitutiva para a segunda-feira, 06/07/2015, gostaria que considerassem a seguinte alteração: prova substitutiva de BAC024 no dia 08/07/2015 às 13h30 na sala de aula s7p Campus José Alencar Anexo III.

Desde já agradeço a atenção,

prof. Gustavo Salgado
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"prova", "substituta", "alteração"})

	//Washington Batista Vieira
	user = posters[7]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Washington Batista Vieira
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[7]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `CASO NÃO HAJA NENHUM IMPEDIMENTO, A DATA DA ÚLTIMA PROVA SERÁ MUDADA PARA O DIA 23/11.
====
DEVIDO À REUNIÃO DO CONSUNI NÃO HAVERÁ AULA HOJE ÀS 17H40.
ESTAREI NA SALA DE AULA A PARTIR DAS 19H30 E FALAREI SOBRE A RESOLUÇÃO DA TERCEIRA PROVA
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"data", "prova", "alterada"})

	//Walter Aoiama Nagai
	user = posters[8]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da L.O.T.S.
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[8]
	post = &Post{User: user,
		Title:       "site",
		Subtitle:    "de programação básica",
		Description: "acessem o site http://bac004.wikidot.com/start",
		Text: `O site http://bac004.wikidot.com/start tem como ideia possibilitar o desenvolvimento de raciocínio lógico para estabelecer soluções computacionais em um dado problema, empregando técnicas de desenvolvimento de programas corretos e bem estruturados, direcionando estas soluções para codificação em linguagem C/C++.
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"site", "bac004", "c++"})

	//Cai&pira
	user = posters[9]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da republica Cai&pira
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[9]
	post = &Post{User: user,
		Title:       "HELLOGREEN",
		Subtitle:    "Festa open bar",
		Description: "atrações confirmadas",
		Text: `****ATRAÇÃO CONFIRMADA****
Savim ProggY com o melhor do Progressive Trance, para travar o nintendo!!!
>>FACEBOOK: https://www.facebook.com/djsaviovilarino/
>>SOUNDCLOUD: https://soundcloud.com/savimvilari…/savin-under-the-same-sun
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"HELLOGREEN", "festa", "atrações"})

	//Meteolate
	user = posters[10]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da republica Meteolate
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[10]
	post = &Post{User: user,
		Title:       "Jumpira",
		Subtitle:    "Festa noturna",
		Description: "Menos de 2 semanas para o maior e melhor evento universitário da região",
		Text: `Menos de 2 semanas para o maior e melhor evento universitário da região!
1º Lote: R$ 60,00
Não deixe para última hora, você pode adquirir seu pacote com as atléticas participantes ou pela internet.
www.ingressonarede.com/site/index.php/ingressos/esportes
Os ingressos avulsos para festa noturna também já estão a venda em João Monlevade (UPTIME), em Itabira (República Meteolate) e pela internet (Ingressos na Rede).
www.ingressonarede.com/site/index.php/ingressos/festas
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"festa", "jumpira", "ingresso"})

	//DNA
	user = posters[11]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da republica DNA
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[11]
	post = &Post{User: user,
		Title:       "VAGAS",
		Subtitle:    "para bixos",
		Description: "DNA está recrutando bixos, com vagas fixas e temporarias.",
		Text: `##### BIXOS SEJAM BEM VINDOS ####
A Rep.DNA, uma das mais antigas e tradicionais, está recrutando bixos, com varias vagas fixas e temporárias, inclusive com quarto individual, tudo a um preço muito acessível, venha conhecer seus melhores amigos, ter seus piores porres, a seja membro da FAMÍLIA DNA, tudo isso em uma casa super confortável e um ótimo clima.
Contato: Rannier Almeida 12 997336076.
Felipe Conegero 15 99754-2912
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"vagas", "fixas", "temporaria"})

	//Amazonas
	user = posters[12]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da republica Amazonas
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[12]
	post = &Post{User: user,
		Title:       "Festa",
		Subtitle:    "Open bar",
		Description: "Festa open bar com patrocionio BOMBAR",
		Text: `Bora Galeraaaaa!
			OPEN BAR dia 30/09 com NA IDEIA!
			*****INGRESSOS*****
			*BOMBAR
			*Moradores da Rep. Amazonas
			*Gabriel Andrade 98932-3968
			*Rafael Beletable 98622-0612
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"festa", "bombar", "openbar"})

	//Bar do João
	user = posters[13]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste do Bar do João
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[13]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `*****************************************
				...HOJE - " BOTECO" - HOJE...
				*****************************************
				*****************************************
				****** Nas noites de Quinta *******
				*****************************************
				............. ATRAÇÃO ..................
				*****************************************
				.......... Arthur Domingues ..........
				*****************************************
				*****************************************
				...Super Descontos para Você...
				*****************************************
				.............. CONFIRA ....................
				******************************************
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"sertanejo", "festa", "desconto"})

	//Filé di gato
	user = posters[14]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste do Filé di gato
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[14]
	post = &Post{User: user,
		Title:       "Sorteio",
		Subtitle:    "resultados",
		Description: "Resultados dos sorteios para o show da banda Mulher de Banda",
		Text: `Parabéns, Stefania Chaves, Helena Lage Muzzi Cabral, Gabi Azevedo e Carlos Augusto Rocha Linhares. Cada um de vocês ganharam 1 ingresso + 1 combo de bebidas para curtirem a noite hoje com Mulher de Banda no Filé Di Gato! Para os ganhadores, favor entrar em contato através de mensagem inbox do Grupo DeFato e enviar seus dados pessoais para autorizarmos o prêmio. DeFato e Filé di Gato agradecem à todos os participantes! ;)
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"sorteio", "mulherdebanda", "ingresso"})

	//Bombar
	user = posters[15]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste do Bombar
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[15]
	post = &Post{User: user,
		Title:       "Cerveja",
		Subtitle:    " Ashby IPA",
		Description: "Nosso primeiro anúncio",
		Text: `Chegaram as cervejas Ashby IPA. É sucesso na certa!!!
Venham logo...
Loja Ashby Itabira.
#Harmonização
#Cerveja
#IPA
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"cerveja", "ashby", "bar"})

	user = posters[16]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Appricci
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[16]
	post = &Post{User: user,
		Title:       "Vaga para garçom",
		Subtitle:    "Horário noturno",
		Description: "Oferta de vaga de emprego para garçom ou garçonete com experiência no período noturno.",
		Text: `Precisa-se de Garçom ou Garçonete com experiencia para trabalhar em horário noturno. Os interessados deverão comparecer no estabelecimento de posse do seu curriculum de 08:00 as 11:00 hs da manhã.
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"vaga", "emprego", "noturno"})

	user = posters[17]
	post = &Post{User: user,
		Title:       "Apresentação",
		Subtitle:    "Primeiro anúncio",
		Description: "Nosso primeiro anúncio",
		Text: `# Teste da Família Pires
		Esperamos vocês visitando o site e conferindo as nossas novidades!
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"teste", "apresentacao", "itabirapp"})

	user = posters[17]
	post = &Post{User: user,
		Title:       "Ofertas do feriado",
		Subtitle:    "Produtos em oferta",
		Description: "",
		Text: `SEU FERIADO COMEÇA AQUI!
OFERTAS VÁLIDAS DE 06 A 11/10 PARA AS LOJAS DE ITABIRA E SANTA MARIA DE ITABIRA.
Somos uma empresa com atitude sustentável. A partir de agora só trabalhamos com anúncios virtuais.
![](https://scontent-gru2-1.xx.fbcdn.net/v/t1.0-9/14563402_1980800432021651_1286242617984150004_n.jpg?oh=368f9f427b7e28820555cf92bc5f135d&oe=58CFFA99)
![](https://scontent-gru2-1.xx.fbcdn.net/v/t1.0-9/14463296_1980800428688318_3169949579455554512_n.jpg?oh=ecf57655600999b6ce968281eb440a20&oe=58CA4EB0)
![](https://fbcdn-sphotos-e-a.akamaihd.net/hphotos-ak-xpt1/v/t1.0-9/14522717_1980800425354985_5281292350574616990_n.jpg?oh=479cd75e7dff83fab850b6d55d1ac41d&oe=588BBE7B&__gda__=1486877435_a905fb1bf828fc37c2bda832b2f31e96)
		`}
	o.Insert(post)
	AppendTagsForPost(post, []string{user.NameIdTag, user.Institution_Tag,
		"padaria", "oferta", "cerveja"})

}
