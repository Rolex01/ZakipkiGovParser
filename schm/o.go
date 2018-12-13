package schm

type account struct {
	BankAddress string `xml:"bankAddress"` /* Адрес кредитной организации */
	BankName    string `xml:"bankName"`    /* Наименование кредитной организации */
	//	Bik             bikType               `xml:"bik"`             /* БИК */
	CorrAccount string `xml:"corrAccount"` /* Корреспондентский счет */
	//PaymentAccount settlementAccountType `xml:"paymentAccount"` /* Расчетный счет */
	//PersonalAccount personalAccountType   `xml:"personalAccount"` /* Лицевой счет */
}

type admissionResults struct {
	AdmissionResult struct {
		ProtocolCommissionMember commissionMemberType  `xml:"protocolCommissionMember"`
		Admitted                 bool                  `xml:"admitted"`
		AppRejectedReason        appRejectedReasonType `xml:"appRejectedReason"`
	} `xml:"admissionResult"`
}

type applicationFeaturesCorrespondence struct {
	Compatible          bool                    `xml:"compatible"`          /* Флаг соответствия */
	NotificationFeature notificationFeatureType `xml:"notificationFeature"` /* Особенность размещения заказа */
}

type appRejectedReasonType struct {
	NsiRejectReason struct {
		Id     int64  `xml:"id"`
		Reason string `xml:"reason"`
	} `xml:"nsiRejectReason"`
	Explanation string `xml:"explanation"` /* Объяснение */
}

type auctionItemsType struct {
	AuctionItem struct {
		Sid         int64   `xml:"sid"`
		Description string  `xml:"description"`
		Price       float64 `xml:"price"`
	} `xml:"auctionItem"`
}

type bidType struct {
	Price                float64 `xml:"price"`                /* Предложение цены */
	Date                 string  `xml:"date"`                 /* Дата и время подачи ценового предложения */
	IncreaseInitialPrice bool    `xml:"increaseInitialPrice"` /* Признак ценового предложения на повышение начальной(максимальной) цены контракта */
}

type childrenCriteriaType struct {
	Id             int64   `xml:"id"`             /* Идентификатор подкритерия */
	Name           string  `xml:"name"`           /* Наименование подкритерия */
	CriterionValue float64 `xml:"criterionValue"` /* Максимальное количество баллов */
	EvalValue      string  `xml:"evalValue"`      /* Дополнительная информация
	или
	Содержание и порядок оценки по подкритерию */
}

type clarificationRequestType struct {
	NotificationNumber string `xml:"notificationNumber"` /* Номер извещения */
	RegNumber          string `xml:"regNumber"`          /* Номер запроса на разъяснение */
	RegDate            string `xml:"regDate"`            /* Дата поступления запроса на разъяснение */
	Topic              string `xml:"topic"`              /* Тема запроса на разъяснение /
	краткое описание запроса на разъяснение */
	Participant struct {
		Name  string `xml:"name"`
		Email string `xml:"email"`
	} `xml:"participant"`
	DocumentMetas documentList `xml:"documentMetas"` /* Информация о прикрепленных документах */
}

type clarificationType struct {
	Id                 int64  `xml:"id"`                 /* Идентификатор объекта ЕИС */
	NotificationNumber string `xml:"notificationNumber"` /* Номер извещения о торгах, запросе котировок */
	RequestRegNumber   string `xml:"requestRegNumber"`   /* Номер запроса на разъяснение */
	RegNumber          string `xml:"regNumber"`          /* Номер разъяснения */
	VersionNumber      int    `xml:"versionNumber"`      /* Номер редакции */
	CreateDate         string `xml:"createDate"`         /* Дата создания */
	PublishDate        string `xml:"publishDate"`        /* Дата публикации */
	Href               string `xml:"href"`               /* Гиперссылка на опубликованное разъяснение */
	Question           string `xml:"question"`           /* Тема запроса на разъяснение */
	Topic              string `xml:"topic"`              /* Тема разъяснения /
	краткое описание разъяснения */
	DocumentMetas documentList `xml:"documentMetas"` /* Информация о прикрепленных документах */
}

type commissionMemberType struct {
	Id   int64  `xml:"id"`   /* Идентификатор члена комиссии */
	Name string `xml:"name"` /* ФИО */
	Role struct {
		RoleId int64  `xml:"roleId"`
		Name   string `xml:"name"`
	} `xml:"role"`
}

type commissionType struct {
	RegNumber         int    `xml:"regNumber"`      /* Номер в реестре комиссий */
	CommissionName    string `xml:"commissionName"` /* Название комиссии */
	CommissionMembers struct {
		CommissionMember string `xml:"commissionMember"`
	} `xml:"commissionMembers"`
	Owner  organizationRef `xml:"owner"`  /* Организация, размещающая заказ */
	Actual bool            `xml:"actual"` /* Признак актуальности */
}

type commonAttributesType struct {
}

type contactInfoType struct {
	OrgName        string            `xml:"orgName"`        /* Наименование организации, размещающей заказа */
	OrgFactAddress string            `xml:"orgFactAddress"` /* Адрес местонахождения организации */
	OrgPostAddress string            `xml:"orgPostAddress"` /* Почтовый адрес организации */
	ContactPerson  contactPersonType `xml:"contactPerson"`  /* Контактное лицо */
	ContactEMail   string            `xml:"contactEMail"`   /* e-mail адрес контактного лица */
	ContactPhone   string            `xml:"contactPhone"`   /* Телефон контактного лица */
	ContactFax     string            `xml:"contactFax"`     /* Факс контактного лица */
}

type contactPersonType struct {
	LastName   string `xml:"lastName"`   /* Фамилия */
	FirstName  string `xml:"firstName"`  /* Имя */
	MiddleName string `xml:"middleName"` /* Отчество */
}

type contractConditionAttributesType struct {
}

type contractSignType struct {
	Id         int64  `xml:"id"`       /* Идентификатор документа ЕИС */
	Number     string `xml:"number"`   /* Номер контракта */
	SignDate   string `xml:"signDate"` /* Дата заключения контракта */
	Foundation struct {
		Order `xml:"order"`
	} `xml:"foundation"`
	Customer     organizationRef `xml:"customer"`     /* Заказчик */
	ProtocolDate string          `xml:"protocolDate"` /* Дата поведения итогов электронного аукциона */
	Suppliers    struct {
		Supplier participantType `xml:"supplier"`
	} `xml:"suppliers"`
	Scan          documentList `xml:"scan"`          /* Контракт в электронной форме */
	DocumentMetas documentList `xml:"documentMetas"` /* Информация о прикрепленных документах */
}

type countryType struct {
	CountryCode     string `xml:"countryCode"`     /* Цифровой код страны */
	CountryFullName string `xml:"countryFullName"` /* Полное наименование страны */
}

type criterion struct {
	NsiEvalCriterion evalCriterion        `xml:"nsiEvalCriterion"` /* Критерий оценки заявки */
	CriterionValue   float64              `xml:"criterionValue"`   /* Значимость критерия */
	ChildrenCriteria childrenCriteriaType `xml:"childrenCriteria"` /* Список подкритериев для данного критерия. Критерий должен быть критерием экспертной оценки, в противном случае список должен быть пуст */
	EvalValue        string               `xml:"evalValue"`        /* Для случая оценки заявок в ручном режиме: текстовое поле «Содержание и порядок оценки по критерию»

	Для случая оценки заявок в соответствии с Порядком оценки заявок (автоматический режим): большое текстовое поле «Дополнительная информация», «Предмет оценки» и т.п. в зависимости от типа критерия. */
}

type criterionRef struct {
	Id int64 `xml:"id"` /* Идентификатор критерия (nsiEvalCriterion/id)
	или
	Идентификатор подкритерия (childrenCriteria/id) */
	CriterionType criterion_criterionType `xml:"criterionType"` /* Дискриминатор, признак, является ли объект критерием или подкритерием

	CR - критерий;
	SC - подкритерий */
}

type currencyType struct {
	Code string `xml:"code"` /* Код валюты */
	Name string `xml:"name"` /* Наименование валюты */
}

type docReqType struct {
	DocumentRequirement `xml:"documentRequirement"` /*  */
}

type document struct {
	Sid            int64  `xml:"sid"`            /* Уникальный идентификатор в ЕИС */
	FileName       string `xml:"fileName"`       /* Имя файла */
	DocDescription string `xml:"docDescription"` /* Описание прикрепляемого документа */
	CryptoSigns    struct {
		Signature string `xml:"signature"`
	} `xml:"cryptoSigns"`
}

type documentList struct {
	DocumentMeta document `xml:"documentMeta"` /*  */
}

type documentRequirement struct {
	Sid           int64  `xml:"sid"`           /* Уникальный идентификатор в ЕИС */
	OrdinalNumber int    `xml:"ordinalNumber"` /* Порядковый номер */
	ReqValue      string `xml:"reqValue"`      /* Текст требования
	ЭА:Требование о наличии документа установлено на основании нормативного акта */
}

type EPType struct {
	Code string `xml:"code"` /* Кодовое наименование ЭП */
	Name string `xml:"name"` /* Название электронной площадки */
	Url  string `xml:"url"`  /* Адрес электронной площадки */
}

type evalCriterion struct {
	Id    int64  `xml:"id"`    /* Идентификатор */
	Name  string `xml:"name"`  /* Название критерия */
	Order int    `xml:"order"` /* Порядковый номер */
}

type evaluationResult struct {
	CommissionMemberId int64   `xml:"commissionMemberId"` /* Идентификатор члена комиссии */
	EvalResult         float64 `xml:"evalResult"`         /* поле ввода «Оценка», которое становится доступно при установленном флаге «Формирование оценки по каждому члену комиссии» */
}

type guaranteeAppType struct {
	Procedure         string                `xml:"procedure"`         /* Срок и порядок внесения платы */
	SettlementAccount settlementAccountType `xml:"settlementAccount"` /* Номер расчётного счёта внесения платы */
	PersonalAccount   personalAccountType   `xml:"personalAccount"`   /* Номер лицевого счёта внесения платы */
	Bik               bikType               `xml:"bik"`               /* БИК */
}

type guaranteeContractType struct {
	Procedure         string                `xml:"procedure"`         /* Срок и порядок внесения платы */
	SettlementAccount settlementAccountType `xml:"settlementAccount"` /* Номер расчётного счёта внесения платы */
	PersonalAccount   personalAccountType   `xml:"personalAccount"`   /* Номер лицевого счёта внесения платы */
	Bik               bikType               `xml:"bik"`               /* БИК */
	IsBail            bool                  `xml:"isBail"`            /* Установлена ли плата (средства под залог) */
}

type guaranteeType struct {
	Procedure         string                `xml:"procedure"`         /* Срок и порядок внесения платы */
	SettlementAccount settlementAccountType `xml:"settlementAccount"` /* Номер расчётного счёта внесения платы */
	PersonalAccount   personalAccountType   `xml:"personalAccount"`   /* Номер лицевого счёта внесения платы */
	Bik               bikType               `xml:"bik"`               /* БИК */
	Amount            float64               `xml:"amount"`            /* Размер платы (в деньгах) */
	Currency          currencyType          `xml:"currency"`          /* Валюта, в которой установлена плата за документацию */
}

type hrefType struct {
	Open    urlType `xml:"open"`    /* Гиперссылка на открытую часть */
	Private urlType `xml:"private"` /* Гиперссылка на личный кабинет */
}

type KBKType struct {
	Kbk1 string `xml:"kbk1"` /* КБК. Часть 1
	Глава. */
	Kbk2 string `xml:"kbk2"` /* КБК. Часть 2
	Раздел/подраздел */
	Kbk3 string `xml:"kbk3"` /* КБК. Часть 3
	Целевая статья */
	Kbk4 string `xml:"kbk4"` /* КБК. Часть 4
	Вид расходов */
	Kbk5 string `xml:"kbk5"` /* КБК. Часть 5
	КОСГУ */
}

type lotRefType struct {
	Currency currencyType `xml:"currency"` /* Валюта */
	Subject  string       `xml:"subject"`  /* Предмет контракта */
}

type modificationType struct {
	InitiativeType modification_initiativeType `xml:"initiativeType"` /* По чей инициативе производится изменение:

	Z - Решение заказчика, уполномоченного органа;
	U - Предписание органа, уполномоченного на осуществление контроля;
	US - Предписание органа, уполномоченного на осуществление контроля о переносе сроков аукциона без отмены протокола рассмотрения заявок;
	S - Решение суда */
	ModificationDate string                     `xml:"modificationDate"` /* Дата решения по изменению */
	Info             string                     `xml:"info"`             /* Краткая информация об изменении */
	AuthorityType    modification_authorityType `xml:"authorityType"`    /* Вид органа, уполномоченного на осуществление контроля

	FA - Федеральная антимонопольная служба;
	FO - Федеральная служба по оборонному заказу;
	S - Орган исполнительной власти субъекта РФ;
	M - Орган местного самоуправления муниципального района, городского округа */
	AuthorityName         string `xml:"authorityName"`  /* Наименование органа, уполномоченного на осуществление контроля */
	CourtName             string `xml:"courtName"`      /* Наименование суда, принявшего решение */
	CourtDesNumber        string `xml:"courtDesNumber"` /* Номер решения суда */
	DesNumber             string `xml:"desNumber"`      /* № решения комиссии, инспекции контролирующего органа */
	AdditionalInfo        string `xml:"additionalInfo"` /* Дополнительная информация */
	AuthorityPrescription struct {
		CheckResultNumber  checkResultNumberType  `xml:"checkResultNumber"`
		PrescriptionNumber prescriptionNumberType `xml:"prescriptionNumber"`
	} `xml:"authorityPrescription"`
	DocumentMetas documentList `xml:"documentMetas"` /* Информация о прикрепленных документах */
}

type notificationCancelType struct {
	Id                 int64  `xml:"id"`                 /* Идентификатор документа ЕИС */
	NotificationNumber string `xml:"notificationNumber"` /* Номер извещения */
	CreateDate         string `xml:"createDate"`         /* Дата создания */
	PublishDate        string `xml:"publishDate"`        /* Дата публикации
	Планируемая или фактическая */
	VersionNumber int `xml:"versionNumber"` /* Номер редакции */
	Lots          struct {
		Lot `xml:"lot"`
	} `xml:"lots"`
	PrintForm                 document                          `xml:"printForm"`     /* Печатная форма извещения */
	DocumentMetas             documentList                      `xml:"documentMetas"` /* Информация о прикрепленных документах */
	Modification              modificationType                  `xml:"modification"`  /* Основание внесения изменений */
	NotificationCancelFailure `xml:"notificationCancelFailure"` /* Сведения об отмене извещения об отказе */
}

type notificationEFType struct {
	BasedOn                notificationType
	NotificationCommission struct {
		P1Date string `xml:"p1Date"` /* Дата и время окончания срока подачи заявок */
		P2Date string `xml:"p2Date"` /* Дата окончания срока рассмотрения первых частей заявок */
		P3Date string `xml:"p3Date"` /* Дата проведения открытого аукциона в электронной форме */
	} `xml:"notificationCommission"`
	Lots struct {
		Lot `xml:"lot"` /*  */
	} `xml:"lots"`
	EP EPType `xml:"EP"` /* Электронная площадка */
}

type notificationFeatureType struct {
	PrefValue        float64 `xml:"prefValue"` /* Величина (преимущества) */
	PlacementFeature struct {
		Code int64  `xml:"code"`
		Name string `xml:"name"`
	} `xml:"placementFeature"`
}

type notificationOKType struct {
	BasedOn               notificationType
	NotificationPlacement struct {
		AdditionalInfo       string                       `xml:"additionalInfo"`    /* Дополнительная информация */
		GuaranteeApp         guaranteeAppType             `xml:"guaranteeApp"`      /* Обязательство обеспечения заявки */
		GuaranteeContract    guaranteeContractType        `xml:"guaranteeContract"` /* Обязательство обеспечения исполнения контракта */
		NotificationFeatures `xml:"notificationFeatures"` /* Особенности размещения заказа */
	} `xml:"notificationPlacement"`
	CompetitiveDocumentProvisioning struct {
		DeliveryTerm      string            `xml:"deliveryTerm"`      /* Срок предоставления документации. С */
		DeliveryTerm2     string            `xml:"deliveryTerm2"`     /* Срок предоставления документации. По */
		DeliveryPlace     string            `xml:"deliveryPlace"`     /* Место предоставления документации */
		DeliveryProcedure string            `xml:"deliveryProcedure"` /* Порядок предоставления документации */
		Www               string            `xml:"www"`               /* Официальный сайт, на котором размещена документация */
		Guarantee         `xml:"guarantee"` /* Плата за документацию */
	} `xml:"competitiveDocumentProvisioning"`
	NotificationCommission struct {
		P1Date  string `xml:"p1Date"`  /* Дата и время вскрытия конвертов с заявками */
		P1Place string `xml:"p1Place"` /* Место вскрытия конвертов с заявками на участие в конкурсе */
		P2Date  string `xml:"p2Date"`  /* Дата рассмотрения заявок */
		P2Place string `xml:"p2Place"` /* Место рассмотрения заявок */
		P3Date  string `xml:"p3Date"`  /* Дата подведения итогов */
		P3Place string `xml:"p3Place"` /* Место подведения итогов */
	} `xml:"notificationCommission"`
	Lots struct {
		Lot `xml:"lot"` /*  */
	} `xml:"lots"`
}

type notificationPlacerChangeType struct {
	Id                 int64               `xml:"id"`                 /* Идентификатор документа ЕИС */
	NotificationNumber string              `xml:"notificationNumber"` /* Номер извещения */
	VersionNumber      int                 `xml:"versionNumber"`      /* Номер редакции */
	CreateDate         string              `xml:"createDate"`         /* Дата создания */
	PublishDate        string              `xml:"publishDate"`        /* Дата публикации */
	ChangeType         orgPlacerChangeType `xml:"changeType"`         /* Основаниe для изменения организации, размещающей заказ

	C - Прекращение действия права уполномоченного органа на размещение заказов для заказчика
	R - Реорганизация заказчика или уполномоченного органа, размещающего заказ */
	PlacerChange struct {
		CurrentPlacer organizationRef `xml:"currentPlacer"`
		NewPlacer     organizationRef `xml:"newPlacer"`
	} `xml:"placerChange"`
	InitiatorChange struct {
		CurrentInitiator organizationRef `xml:"currentInitiator"`
		NewInitiator     organizationRef `xml:"newInitiator"`
	} `xml:"initiatorChange"`
	Comment       string       `xml:"comment"`       /* Дополнительная информация */
	Href          string       `xml:"href"`          /* Гиперссылка на опубликованное разъяснение */
	PrintForm     document     `xml:"printForm"`     /* Печатная форма уведомления */
	DocumentMetas documentList `xml:"documentMetas"` /* Информация о прикрепленных документах */
}

type notificationPOType struct {
	BasedOn                notificationType
	NotificationCommission struct {
		P1Date  string `xml:"p1Date"`  /* Дата начала срока подачи заявок на участие */
		P2Date  string `xml:"p2Date"`  /* Дата окончания срока подачи заявок на участие */
		P1Place string `xml:"p1Place"` /* Место подачи заявок на участие */
		P2Place string `xml:"p2Place"` /* Порядок подачи заявок на участие */
		P3Date  string `xml:"p3Date"`  /* Дата и время проведения предварительного отбора */
		P3Place string `xml:"p3Place"` /* Место проведения предварительного отбора */
	} `xml:"notificationCommission"`
	Lots struct {
		Lot `xml:"lot"` /*  */
	} `xml:"lots"`
}

type notificationSZType struct {
	BasedOn                notificationType
	NotificationCommission struct {
		P1Date string `xml:"p1Date"` /* Срок предоставления предложений */
	} `xml:"notificationCommission"`
	Lots struct {
		Lot `xml:"lot"` /*  */
	} `xml:"lots"`
}

type notificationType struct {
	Id                           int64  `xml:"id"`                           /* Идентификатор документа ЕИС */
	NotificationNumber           string `xml:"notificationNumber"`           /* Реестровый номер извещения */
	FoundationNotificationNumber string `xml:"foundationNotificationNumber"` /* Реестровый номер родительского извещения, используется при выделении лотов в отдельную процедуру размещения заказа */
	VersionNumber                int    `xml:"versionNumber"`                /* Номер редакции */
	CreateDate                   string `xml:"createDate"`                   /* Дата создания */
	PlacingWay                   struct {
		Code string `xml:"code"`
		Name string `xml:"name"`
	} `xml:"placingWay"`
	OrderName string `xml:"orderName"` /* Наименование заказа */
	Order     struct {
		Initiator        organizationRef  `xml:"initiator"`
		InitiatorOrgRole organizationRole `xml:"initiatorOrgRole"`
		Placer           organizationRef  `xml:"placer"`
		PlacerOrgType    organizationType `xml:"placerOrgType"`
	} `xml:"order"`
	ContactInfo   contactInfoType `xml:"contactInfo"`   /* Контактная организация */
	PrintForm     document        `xml:"printForm"`     /* Печатная форма извещения */
	DocumentMetas documentList    `xml:"documentMetas"` /* Информация о прикрепленных документах */
	PublishDate   string          `xml:"publishDate"`   /* Дата публикации
	Планируемая или фактическая. При выгрузке дата публикации первой версии */
	PublishVersionDate string `xml:"publishVersionDate"` /* Дата публикации
	текущей версии */
	Modification modificationType `xml:"modification"` /* Основание внесения изменений */
	Href         string           `xml:"href"`         /* Гиперссылка на опубликованное извещение */
}

type notificationZKType struct {
	BasedOn                         notificationType
	CompetitiveDocumentProvisioning struct {
		Www string `xml:"www"` /* Официальный сайт, на котором размещена документация */
	} `xml:"competitiveDocumentProvisioning"`
	NotificationCommission struct {
		P1Date   string `xml:"p1Date"`   /* Дата и время начала подачи заявок */
		P2Date   string `xml:"p2Date"`   /* Дата и время окончания подачи заявок */
		P1Place  string `xml:"p1Place"`  /* Форма заявки */
		P2Place  string `xml:"p2Place"`  /* Место подачи котировочных заявок */
		SignTerm int64  `xml:"signTerm"` /* Срок подписания победителем контракта
		в днях */
	} `xml:"notificationCommission"`
	Lots struct {
		Lot `xml:"lot"` /*  */
	} `xml:"lots"`
}

type OKEIType struct {
	Code string `xml:"code"` /* Код */
	Name string `xml:"name"` /* Наименование */
}

type okopfType struct {
	Code         string `xml:"code"`         /* Код */
	SingularName string `xml:"singularName"` /* Наименование в единственном числе */
}

type organizationRef struct {
	RegNum   string `xml:"regNum"`   /* Реестровый номер в СПЗ */
	FullName string `xml:"fullName"` /* Полное наименование */
}

type participantType struct {
	ParticipantType participantTypeType `xml:"participantType"` /* Тип участника:
	P - Физическое лицо РФ,
	PF - Физическое лицо иностранного государства,
	U - Юридическое лицо РФ,
	UF - Юридическое лицо иностранного государства */
	Inn               innType              `xml:"inn"`               /* ИНН */
	Kpp               kppType              `xml:"kpp"`               /* КПП */
	OrganizationForm  organizationFormType `xml:"organizationForm"`  /* Тип организационно - правовой формы.Устарело. Не применяется. */
	LegalForm         okopfType            `xml:"legalForm"`         /* Организационно-правовая форма организации в ОКОПФ */
	IdNumber          string               `xml:"idNumber"`          /* Идентификационный номер для ParticipantType.UF */
	IdNumberExtension string               `xml:"idNumberExtension"` /* Дополнительный идентификационный номер для ParticipantType.UF */
	OrganizationName  string               `xml:"organizationName"`  /* Имя организации */
	Country           countryType          `xml:"country"`           /* Код страны в ОКСМ */
	FactualAddress    string               `xml:"factualAddress"`    /* адрес */
	PostAddress       string               `xml:"postAddress"`       /* Почтовый адрес. */
	ContactInfo       contactPersonType    `xml:"contactInfo"`       /* Контактная информация */
	ContactEMail      string               `xml:"contactEMail"`      /* e-mail адрес контактного лица */
	ContactPhone      string               `xml:"contactPhone"`      /* Телефон контактного лица */
	ContactFax        string               `xml:"contactFax"`        /* Факс контактного лица */
	AdditionalInfo    string               `xml:"additionalInfo"`    /* Дополнительная информация */
	Status            string               `xml:"status"`            /* Статусы поставщика (исполнителя, подрядчика)
	– 01, субъект малого предпринимательства
	– 02, учреждения уголовно-исправительной системы
	– 03, общероссийские общественные организации инвалидов */
}

type placementResultType struct {
	NotificationNumber string             `xml:"notificationNumber"` /* Номер извещения о проведении */
	ProtocolNumber     protocolNumberType `xml:"protocolNumber"`     /* Номер протокола */
	VersionNumber      int                `xml:"versionNumber"`      /* Номер редакции */
	Status             statusType         `xml:"status"`             /* Статус результата

	V - Действителен
	N - Недействителен */
	CreateDate           string             `xml:"createDate"`           /* Дата создания */
	LastUpdateDate       string             `xml:"lastUpdateDate"`       /* Дата последнего изменения */
	LotNumber            int                `xml:"lotNumber"`            /* Порядковый номер лота */
	OrderPlacementCancel bool               `xml:"orderPlacementCancel"` /* Признак несостоявшегося размещения заказа */
	ChangePossible       changePossibleType `xml:"changePossible"`       /* Признак возможности изменения результата

	N - Нет;
	O - Другие участники;
	D - Любые комбинации */
	RepeatedPlacement repeatedPlacementType `xml:"repeatedPlacement"` /* Повторное размещение заказа

	Y - Да;
	N - Нет;
	P - Возможно */
	ContractWithParticipant       bool                              `xml:"contractWithParticipant"`       /* Признак заключения контракта с участником */
	ContractWithSingleApplication contractWithSingleApplicationType `xml:"contractWithSingleApplication"` /* Признак заключения контракта по единственной заявки

	Y - Да;
	N - Нет;
	P - Возможно */
	Applications struct {
		Application `xml:"application"`
	} `xml:"applications"`
}

type productsType struct {
	Product productType `xml:"product"` /*  */
}

type productType struct {
	Code string `xml:"code"` /* Код товара, работы или услуги */
	Name string `xml:"name"` /* Наименование товара, работы или услуги */
}

type protocolCancelType struct {
	NotificationNumber string             `xml:"notificationNumber"` /* Номер извещения о проведении */
	ProtocolNumber     protocolNumberType `xml:"protocolNumber"`     /* Номер протокола */
	CreateDate         string             `xml:"createDate"`         /* Дата создания */
	PublishDate        string             `xml:"publishDate"`        /* Дата публикации */
	ProtocolLots       struct {
		ProtocolLot `xml:"protocolLot"`
	} `xml:"protocolLots"`
	PrintForm     document         `xml:"printForm"`     /* Печатная форма */
	DocumentMetas documentList     `xml:"documentMetas"` /* Метаинформация о прикрепленных документах */
	Modification  modificationType `xml:"modification"`  /* Основание отмены протокола */
	Href          string           `xml:"href"`          /* Гиперссылка на опубликованные сведения об отмене протокола */
}

type protocolEF1Type struct {
	BasedOn      protocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /*  */
	} `xml:"protocolLots"`
}

type protocolEF2Type struct {
	NotificationNumber       string             `xml:"notificationNumber"`       /* Номер извещения о проведении */
	ProtocolNumber           protocolNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber protocolNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     protocolNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	VersionNumber            int                `xml:"versionNumber"`            /* Номер редакции */
	PublishDate              string             `xml:"publishDate"`              /* Дата публикации */
	PrintForm                document           `xml:"printForm"`                /* Печатная форма протокола */
	DocumentMetas            documentList       `xml:"documentMetas"`            /* Метаинформация о прикрепленных документах */
	Modifications            modificationType   `xml:"modifications"`            /* Изменения для данной редакции протокола */
	Href                     string             `xml:"href"`                     /* Гиперссылка на опубликованный протокол */
	ProtocolLots             struct {
		ProtocolLot `xml:"protocolLot"`
	} `xml:"protocolLots"`
}

type protocolEF3Type struct {
	BasedOn      protocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /*  */
	} `xml:"protocolLots"`
}

type protocolEvasionType struct {
	Id                       int64        `xml:"id"`                       /* Идентификатор объекта ЕИС */
	NotificationNumber       string       `xml:"notificationNumber"`       /* Номер извещения о проведении */
	ProtocolNumber           string       `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber string       `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     string       `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    string       `xml:"place"`                    /* Место составления протокола */
	VersionNumber            int          `xml:"versionNumber"`            /* Номер редакции */
	ProtocolDate             string       `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string       `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string       `xml:"publishDate"`              /* Дата публикации */
	PrintForm                document     `xml:"printForm"`                /* Печатная форма протокола */
	DocumentMetas            documentList `xml:"documentMetas"`            /* Метаинформация о прикрепленных документах */
	Href                     string       `xml:"href"`                     /* Гиперссылка на опубликованный протокол */
	ProtocolLots             struct {
		ProtocolLot `xml:"protocolLot"`
	} `xml:"protocolLots"`
	RefusalFacts struct {
		RefusalFact refusalFact `xml:"refusalFact"`
	} `xml:"refusalFacts"`
	NewProtocolTypeIndication bool            `xml:"newProtocolTypeIndication"` /* Признак нового протокола */
	Customer                  organizationRef `xml:"customer"`                  /* Заказчик */
}

type protocolOK1Type struct {
	Id                       int64              `xml:"id"`                       /* Идентификатор объекта ЕИС */
	NotificationNumber       string             `xml:"notificationNumber"`       /* Номер извещения о проведении */
	ProtocolNumber           protocolNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber protocolNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     protocolNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	VersionNumber            int                `xml:"versionNumber"`            /* Номер редакции */
	Place                    string             `xml:"place"`                    /* Место составления протокола */
	ProtocolDate             string             `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string             `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string             `xml:"publishDate"`              /* Дата публикации */
	Commission               struct {
		RegNumber                 int    `xml:"regNumber"`
		Name                      string `xml:"name"`
		ProtocolCommissionMembers `xml:"protocolCommissionMembers"`
	} `xml:"commission"`
	PrintForm     document         `xml:"printForm"`     /* Печатная форма протокола */
	DocumentMetas documentList     `xml:"documentMetas"` /* Информация о прикрепленных документах */
	Modification  modificationType `xml:"modification"`  /* Основание внесения изменений */
	Href          string           `xml:"href"`          /* Гиперссылка на опубликованный протокол */
	ProtocolLots  struct {
		ProtocolLot `xml:"protocolLot"`
	} `xml:"protocolLots"`
}

type protocolOK2Type struct {
	BasedOn      protocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /*  */
	} `xml:"protocolLots"`
}

type protocolOK3Type struct {
	BasedOn      protocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /*  */
	} `xml:"protocolLots"`
}

type protocolPO1Type struct {
	BasedOn      protocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /*  */
	} `xml:"protocolLots"`
}

type protocolType struct {
	Id                       int64              `xml:"id"`                       /* Идентификатор объекта ЕИС */
	NotificationNumber       string             `xml:"notificationNumber"`       /* Номер извещения о проведении */
	ProtocolNumber           protocolNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber protocolNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     protocolNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	VersionNumber            int                `xml:"versionNumber"`            /* Номер редакции */
	Place                    string             `xml:"place"`                    /* Место составления протокола */
	ProtocolDate             string             `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string             `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string             `xml:"publishDate"`              /* Дата публикации */
	Commission               struct {
		RegNumber                 int    `xml:"regNumber"`
		Name                      string `xml:"name"`
		ProtocolCommissionMembers `xml:"protocolCommissionMembers"`
	} `xml:"commission"`
	PrintForm     document         `xml:"printForm"`     /* Печатная форма протокола */
	DocumentMetas documentList     `xml:"documentMetas"` /* Информация о прикрепленных документах */
	Modification  modificationType `xml:"modification"`  /* Основание внесения изменений */
	Href          string           `xml:"href"`          /* Гиперссылка на опубликованный протокол */
}

type protocolZK1Type struct {
	BasedOn      protocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /*  */
	} `xml:"protocolLots"`
}

type protocolZK5Type struct {
	OrderName string `xml:"orderName"` /* Наименование заказа */
	Order     struct {
		Placer        organizationRef  `xml:"placer"`
		PlacerOrgType organizationType `xml:"placerOrgType"`
	} `xml:"order"`
	ContactInfo  contactInfoType `xml:"contactInfo"` /* Контактная организация */
	Currency     currencyType    `xml:"currency"`    /* Валюта */
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"`
	} `xml:"protocolLots"`
}

type questionType struct {
	NotificationNumber string       `xml:"notificationNumber"` /* Номер извещения о проведении */
	RegNumber          string       `xml:"regNumber"`          /* Номер запроса на разъяснение */
	RegDate            string       `xml:"regDate"`            /* Дата регистрации */
	Subject            string       `xml:"subject"`            /* Тема запроса на разъяснение */
	Href               string       `xml:"href"`               /* Гиперссылка на опубликованный запрос на разъяснение */
	DocumentMetas      documentList `xml:"documentMetas"`      /*  */
}

type refusalFact struct {
	VoucherEntry string                `xml:"voucherEntry"` /* Реквизиты подтверждающих документов */
	Explanation  string                `xml:"explanation"`  /* Пояснения */
	Foundation   refusalFactFoundation `xml:"foundation"`   /* Основание отказа */
}

type refusalFactFoundation struct {
	Id   int64  `xml:"id"`   /* Идентификатор в справочнике */
	Name string `xml:"name"` /* Наименование */
}

type requirementCompliances struct {
	RequirementCompliance struct {
		OrdinalNumber    int                                     `xml:"ordinalNumber"`
		AvailabilityType requirementCompliances_availabilityType `xml:"availabilityType"`
		Reason           string                                  `xml:"reason"`
	} `xml:"requirementCompliance"`
}

type tendePlanInfoType struct {
	PlanNumber         tenderPlanNumberType         `xml:"planNumber"`         /* Реестровый номер плана-графика */
	PlanPositionNumber tenderPlanPositionNumberType `xml:"planPositionNumber"` /* Номер позиции в плане-графике */
}

type timeEFType struct {
	NotificationNumber string `xml:"notificationNumber"` /* Номер извещения о проведении */
	AuctionTime        string `xml:"auctionTime"`        /* Время проведения */
}

type zfcs_attachmentListType struct {
	Attachment zfcs_attachmentType `xml:"attachment"` /*  */
}

type zfcs_attachmentType struct {
	PublishedContentId zfcs_guidType `xml:"publishedContentId"` /* Уникальный идентификатор контента прикрепленного документа на ЕИС  */
	FileName           string        `xml:"fileName"`           /* Имя файла */
	FileSize           string        `xml:"fileSize"`           /* Размер файла */
	DocDescription     string        `xml:"docDescription"`     /* Описание прикрепляемого документа */
	CryptoSigns        struct {
		Signature string `xml:"signature"`
	} `xml:"cryptoSigns"`
}

type zfcs_contactInfoType struct {
	OrgPostAddress zfcs_longTextType      `xml:"orgPostAddress"` /* Почтовый адрес организации */
	OrgFactAddress zfcs_longTextType      `xml:"orgFactAddress"` /* Адрес местонахождения организации */
	ContactPerson  zfcs_contactPersonType `xml:"contactPerson"`  /* Ответственное должностное лицо */
	ContactEMail   zfcs_string            `xml:"contactEMail"`   /* e-mail адрес контактного лица */
	ContactPhone   zfcs_string            `xml:"contactPhone"`   /* Телефон контактного лица */
	ContactFax     zfcs_string            `xml:"contactFax"`     /* Факс контактного лица */
	AddInfo        zfcs_longTextType      `xml:"addInfo"`        /* Дополнительная информация */
}

type zfcs_contactPersonType struct {
	LastName   string `xml:"lastName"`   /* Фамилия */
	FirstName  string `xml:"firstName"`  /* Имя */
	MiddleName string `xml:"middleName"` /* Отчество */
}

type zfcs_abandonedReasonType struct {
	Code       string                       `xml:"code"`       /* Код основания признания процедуры несостоявшейся */
	ObjectName string                       `xml:"objectName"` /* Наименование интеграционного объекта, к которому относится данное основание */
	Name       string                       `xml:"name"`       /* Наименование основания признания процедуры несостоявшейся */
	Type       zfcs_abandonedReasonTypeEnum `xml:"type"`       /* Тип основания:
	OR - По окончании срока подачи заявок подана только одна заявка. Такая заявка признана соответствующей требованиям 44-ФЗ и требованиям, указанным в извещении;
	NR - По окончании срока подачи заявок не подано ни одной заявки;
	OV - По результатам рассмотрения заявок только одна заявка признана соответствующей требованиям ФЗ и требованиям, указанным в извещении;
	NV - Все поданные заявки отклонены;
	OV2 - По результатам рассмотрения вторых частей заявок только одна заявка признана соответствующей требованиям 44-ФЗ и требованиям, указанным в извещении или ни одной заявки не признано соответствующим данным требованиям */
}

type zfcs_publicDiscussionResultType struct {
	PublicDiscussionNum zfcs_publicDiscussionNumType `xml:"publicDiscussionNum"` /* Реестровый номер общественного обсуждения */
	RevisionNumber      zfcs_revisionNumType         `xml:"revisionNumber"`      /* Номер редакции */
	PublishDate         string                       `xml:"publishDate"`         /* Дата публикации общественного обсуждния */
	UpdatePublishDate   string                       `xml:"UpdatePublishDate"`   /* Дата и время последнего обновления */
	Phase               string                       `xml:"phase"`               /* Стадия общественного обсуждения
	Возможные значения:
	S_1 - Этап 1
	S1F - Этап 1 завершен
	S_2 - Этап 2
	FIN - Обсуждение завершено */
	Customer zfcs_purchaseOrganizationType `xml:"customer"` /* Заказчик */
	Stages   struct {
		Stage `xml:"stage"`
	} `xml:"stages"`
	Place       zfcs_longTextType       `xml:"place"`       /* Место проведения публичных слушаний, порядок доступа к участию */
	Date        string                  `xml:"date"`        /* Дата и время проведения публичных слушаний */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Документы */
}

type zfcs_baseRef struct {
	Code string `xml:"code"` /* Код  */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_decisionRef struct {
	Code string `xml:"code"` /* Код  */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_timelineViolationType struct {
	RefId string `xml:"refId"` /* Идентификатор обработанного пакета данных */
}

type zfcs_publicDiscussionPurchaseInfoType struct {
	PlanNumber           zfcs_tenderPlanNumberType     `xml:"planNumber"`           /* Реестровый номер плана-графика */
	PositionNumber       zfcs_longTextMinType          `xml:"positionNumber"`       /* Номер позиции (номер заказа (лота)) в плане-графике */
	PurchaseNumber       zfcs_purchaseNumberType       `xml:"purchaseNumber"`       /* Реестровый номер закупки */
	LotNumber            int                           `xml:"lotNumber"`            /* Номер лота в извещении */
	PlanObjectInfo       zfcs_longTextMinType          `xml:"planObjectInfo"`       /* Наименование предмета контракта на основании ПГ */
	PlacingWay           zfcs_longTextMinType          `xml:"placingWay"`           /* Способ определения поставщика на основании ПГ */
	Year                 zfcs_yearType                 `xml:"year"`                 /* Год плана-графика */
	PlanContractMaxPrice zfcs_moneyPositiveType        `xml:"planContractMaxPrice"` /* Ориентировочная (начальная) максимальная стоимость контракта на основании ПГ */
	Customer             zfcs_purchaseOrganizationType `xml:"customer"`             /* Заказчик */
}

type zfcs_publicDiscussionLargePurchaseType struct {
	ExternalId          zfcs_externalIdType                          `xml:"externalId"`          /* Внешний идентификатор документа */
	Id                  int64                                        `xml:"id"`                  /* Идентификатор документа */
	PublicDiscussionNum zfcs_publicDiscussionNumType                 `xml:"publicDiscussionNum"` /* Реестровый номер общественного обсуждения */
	VersionNumber       zfcs_revisionNumType                         `xml:"versionNumber"`       /* Номер редакции */
	DocPublishDate      string                                       `xml:"docPublishDate"`      /* Дата публикации документа */
	PublishOrg          zfcs_organizationInfoType                    `xml:"publishOrg"`          /* Организация, разместившая информацию */
	Topic               zfcs_longTextMinType                         `xml:"topic"`               /* Тема обсуждения */
	Phase               zfcs_publicDiscussionLargePurchaseStagesEnum `xml:"phase"`               /* Этап общественного обсуждения.
	Возможные значения:
	S1 - Этап 1
	S2 - Этап 2 */
	Purchase zfcs_publicDiscussionPurchaseInfoType `xml:"purchase"` /* Сведения о позиции плана-графика */
	Hearing  struct {
		Date  string               `xml:"date"`
		Place zfcs_longTextMinType `xml:"place"`
	} `xml:"hearing"`
	Decision struct {
		StartDate string `xml:"startDate"`
		EndDate   string `xml:"endDate"`
	} `xml:"decision"`
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Документы */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_publicDiscussionCommentType struct {
	ExternalId            zfcs_externalIdType           `xml:"externalId"`            /* Внешний идентификатор документа */
	Id                    int64                         `xml:"id"`                    /* Идентификатор документа */
	PublicDiscussionNum   zfcs_publicDiscussionNumType  `xml:"publicDiscussionNum"`   /* Реестровый номер общественного обсуждения */
	VersionNumber         zfcs_revisionNumType          `xml:"versionNumber"`         /* Номер редакции */
	DocPublishDate        string                        `xml:"docPublishDate"`        /* Дата и время прохождения премодерации комментария на ЕИС */
	PublicDiscussionFacet zfcs_publicDiscussionFacetRef `xml:"publicDiscussionFacet"` /* Аспект общественного обсуждения */
	Author                struct {
		Name  zfcs_longTextMinType `xml:"name"`
		EMail zfcs_string          `xml:"eMail"`
	} `xml:"author"`
	CommentNumber int                                          `xml:"commentNumber"` /* Номер комментария */
	Comment       zfcs_longTextMinType                         `xml:"comment"`       /* Текст комментария */
	Phase         zfcs_publicDiscussionLargePurchaseStagesEnum `xml:"phase"`         /* Этап общественного обсуждения.
	Возможные значения:
	S1 - Этап 1
	S2 - Этап 2 */
	Purchase    zfcs_publicDiscussionPurchaseInfoType `xml:"purchase"`    /* Сведения о позиции плана-графика, закупки. Элемент не используется в импорте */
	Attachments zfcs_attachmentListType               `xml:"attachments"` /* Документы */
}

type zfcs_publicDiscussionFormType struct {
	ExternalId          zfcs_externalIdType                 `xml:"externalId"`          /* Внешний идентификатор документа */
	Id                  int64                               `xml:"id"`                  /* Идентификатор документа */
	PublicDiscussionNum zfcs_publicDiscussionNumType        `xml:"publicDiscussionNum"` /* Реестровый номер общественного обсуждения */
	VersionNumber       zfcs_revisionNumType                `xml:"versionNumber"`       /* Номер редакции */
	DocPublishDate      string                              `xml:"docPublishDate"`      /* Дата и время отправки анкеты из открытой части  ЕИС */
	FormNumber          zfcs_publicDiscussionFormNumberType `xml:"formNumber"`          /* Реестровый номер анкеты */
	Author              struct {
		Name  zfcs_longTextMinType `xml:"name"`
		EMail zfcs_string          `xml:"eMail"`
	} `xml:"author"`
	Phase zfcs_publicDiscussionLargePurchaseStagesEnum `xml:"phase"` /* Этап общественного обсуждения.
	Возможные значения:
	S1 - Этап 1
	S2 - Этап 2 */
	PublicDiscussionFacets struct {
		PublicDiscussionFacet `xml:"publicDiscussionFacet"`
	} `xml:"publicDiscussionFacets"`
	Purchase zfcs_publicDiscussionPurchaseInfoType `xml:"purchase"` /* Сведения о позиции плана-графика, закупки. Элемент не используется в импорте */
}

type zfcs_publicDiscussionAnswerType struct {
	Id                  int64                                        `xml:"id"`                  /* Идентификатор документа */
	ExternalId          zfcs_externalIdType                          `xml:"externalId"`          /* Внешний идентификатор документа */
	PublicDiscussionNum zfcs_publicDiscussionNumType                 `xml:"publicDiscussionNum"` /* Реестровый номер общественного обсуждения */
	VersionNumber       zfcs_revisionNumType                         `xml:"versionNumber"`       /* Номер редакции */
	DocPublishDate      string                                       `xml:"docPublishDate"`      /* Дата и время публикации ответа на комментарий на ЕИС */
	PublishOrg          zfcs_organizationInfoType                    `xml:"publishOrg"`          /* Организация, разместившая информацию */
	CommentNumber       int                                          `xml:"commentNumber"`       /* Номер комментария */
	Comment             zfcs_longTextMinType                         `xml:"comment"`             /* Текст комментария. Элемент не используется при импорте. */
	AnswerNumber        int                                          `xml:"answerNumber"`        /* Номер ответа */
	Answer              zfcs_longTextMinType                         `xml:"answer"`              /* Текст ответа */
	Phase               zfcs_publicDiscussionLargePurchaseStagesEnum `xml:"phase"`               /* Этап общественного обсуждения.
	Возможные значения:
	S1 - Этап 1
	S2 - Этап 2 */
	Purchase     zfcs_publicDiscussionPurchaseInfoType `xml:"purchase"`     /* Сведения о позиции плана-графика, закупки. Элемент не используется в импорте */
	Attachments  zfcs_attachmentListType               `xml:"attachments"`  /* Документы */
	ExtPrintForm zfcs_extPrintFormType                 `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_publicDiscussionProtocolType struct {
	ExternalId          zfcs_externalIdType                          `xml:"externalId"`          /* Внешний идентификатор документа */
	Id                  int64                                        `xml:"id"`                  /* Идентификатор документа */
	PublicDiscussionNum zfcs_publicDiscussionNumType                 `xml:"publicDiscussionNum"` /* Реестровый номер общественного обсуждения */
	VersionNumber       zfcs_revisionNumType                         `xml:"versionNumber"`       /* Номер редакции */
	DocPublishDate      string                                       `xml:"docPublishDate"`      /* Дата публикации документа */
	PublishOrg          zfcs_organizationInfoType                    `xml:"publishOrg"`          /* Организация, разместившая информацию */
	Topic               zfcs_longTextMinType                         `xml:"topic"`               /* Тема обсуждения. Элемент не используется в импорте */
	Phase               zfcs_publicDiscussionLargePurchaseStagesEnum `xml:"phase"`               /* Этап общественного обсуждения.
	Возможные значения:
	S1 - Этап 1
	S2 - Этап 2 */
	Decision     zfcs_publicDiscussionDecisionRef      `xml:"decision"`     /* Решение общественного обсуждения */
	Foundation   zfcs_publicDiscussionFoundationRef    `xml:"foundation"`   /* Основание решения общественного обсуждения */
	Purchase     zfcs_publicDiscussionPurchaseInfoType `xml:"purchase"`     /* Сведения о позиции плана-графика, закупки. Элемент не используется в импорте */
	Attachments  zfcs_attachmentListType               `xml:"attachments"`  /* Документы */
	ExtPrintForm zfcs_extPrintFormType                 `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_addInfoType struct {
	ExternalId       zfcs_externalIdType         `xml:"externalId"`       /* Внешний идентификатор документа */
	Id               int64                       `xml:"id"`               /* Идентификатор документа */
	RegistryNum      zfcs_purchaseNumberType     `xml:"registryNum"`      /* Реестровый номер */
	PublishOrg       zfcs_organizationInfoType   `xml:"publishOrg"`       /* Организация, разместившая информацию */
	VersionNumber    zfcs_revisionNumType        `xml:"versionNumber"`    /* Номер редакции */
	DocPublishDate   string                      `xml:"docPublishDate"`   /* Дата публикации документа */
	FirstPublishDate string                      `xml:"firstPublishDate"` /* Дата публикации документа первой редакции */
	Href             zfcs_hrefType               `xml:"href"`             /* Гиперссылка на опубликованные сведения */
	PrintForm        zfcs_contract_printFormType `xml:"printForm"`        /* Печатная форма */
	InfoType         zfcs_addInfoTypeEnum        `xml:"infoType"`         /* Тип информации:
	O - Решение заказчика об одностороннем отказе от исполнения контракта;
	B - Информация о непредоставлении участником закупки сведений о выгодоприобретателях;
	S - Информация о непредоставлении участником закупки сведений о субподрядчиках, соисполнителях.
	*/
	Attachments        zfcs_attachmentListType `xml:"attachments"`        /* Документы */
	ModificationReason zfcs_longTextType       `xml:"modificationReason"` /* Причина редактирования информации */
	ExtPrintForm       zfcs_extPrintFormType   `xml:"extPrintForm"`       /* Электронный документ, полученный из внешней системы */
}

type zfcs_addInfoInvalidType struct {
	ExternalId       zfcs_externalIdType         `xml:"externalId"`       /* Внешний идентификатор документа */
	Id               int64                       `xml:"id"`               /* Идентификатор документа */
	RegistryNum      zfcs_purchaseNumberType     `xml:"registryNum"`      /* Реестровый номер */
	PublishOrg       zfcs_organizationInfoType   `xml:"publishOrg"`       /* Организация, разместившая информацию  */
	VersionNumber    zfcs_revisionNumType        `xml:"versionNumber"`    /* Номер редакции дополнительной информации */
	DocPublishDate   string                      `xml:"docPublishDate"`   /* Дата публикации документа */
	FirstPublishDate string                      `xml:"firstPublishDate"` /* Дата публикации документа первой редакции */
	Href             zfcs_hrefType               `xml:"href"`             /* Гиперссылка на опубликованные сведения */
	PrintForm        zfcs_contract_printFormType `xml:"printForm"`        /* Печатная форма */
	InfoType         zfcs_addInfoTypeEnum        `xml:"infoType"`         /* Тип информации:
	O - Решение заказчика об одностороннем отказе от исполнения контракта;
	B - Информация о непредоставлении участником закупки сведений о выгодоприобретателях;
	S - Информация о непредоставлении участником закупки сведений о субподрядчиках, соисполнителях.
	 (для печатной формы) */
	InvalidityInfo struct {
		Date   string               `xml:"date"`
		Reason zfcs_longTextMinType `xml:"reason"`
	} `xml:"invalidityInfo"`
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Документы */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_auditResultType struct {
	Id               int64                       `xml:"id"`               /* Идентификатор документа */
	ExternalId       zfcs_externalIdType         `xml:"externalId"`       /* Внешний идентификатор документа */
	VersionNumber    zfcs_revisionNumType        `xml:"versionNumber"`    /* Номер редакции */
	RegistryNum      zfcs_purchaseNumberType     `xml:"registryNum"`      /* Реестровый номер */
	DocPublishDate   string                      `xml:"docPublishDate"`   /* Дата публикации документа */
	FirstPublishDate string                      `xml:"firstPublishDate"` /* Дата публикации документа первой редакции */
	PublishOrg       zfcs_organizationInfoType   `xml:"publishOrg"`       /* Орган аудита в сфере закупок, размещающий обобщенную информацию */
	Href             zfcs_hrefType               `xml:"href"`             /* Гиперссылка на опубликованные сведения */
	PrintForm        zfcs_contract_printFormType `xml:"printForm"`        /* Печатная форма */
	ExtPrintForm     zfcs_extPrintFormType       `xml:"extPrintForm"`     /* Электронный документ, полученный из внешней системы */
	Type             zfcs_auditResultTypeEnum    `xml:"type"`             /* Вид обобщенной информации:
	E - Результаты экспертно-аналитической деятельности;
	C - Результаты контрольной деятельности;
	I - Результаты информационной деятельности;
	O - Результаты иной деятельности.
	*/
	Name     zfcs_longTextMinType `xml:"name"` /* Наименование обобщенной информации */
	Document struct {
		Type   string `xml:"type"`
		Date   string `xml:"date"`
		Number string `xml:"number"`
	} `xml:"document"`
	Period struct {
		Start string `xml:"start"`
		End   string `xml:"end"`
	} `xml:"period"`
	AddInfo zfcs_longTextMinType `xml:"addInfo"` /* Дополнительная информация */
	Action  struct {
		AuditOrg `xml:"auditOrg"`
		Subjects `xml:"subjects"`
		Objects  `xml:"objects"`
		Period   `xml:"period"`
	} `xml:"action"`
	Attachments        zfcs_attachmentListType `xml:"attachments"`        /* Документы */
	ModificationReason zfcs_longTextMinType    `xml:"modificationReason"` /* Причина редактирования информации */
}

type zfcs_regulationRulesType struct {
	Id             int64                           `xml:"id"`             /* Идентификатор документа ЕИС */
	DocPublishDate string                          `xml:"docPublishDate"` /* Дата публикации документа */
	RegistryNum    zfcs_standardContractNumberType `xml:"registryNum"`    /* Реестровый номер */
	PublishOrg     `xml:"publishOrg"`              /* Организация, разместившая информацию */
	Type           zfcs_regulationRulesTypeEnum    `xml:"type"` /* Тип правил нормирования:
	1 - Общие правила нормирования в сфере закупок;
	2 - Общие требования к правовым актам по нормированию;
	3 - Общие требования к отдельным видам товаров, работ, услуг;
	4 - Общие требования к определению нормативных затрат;
	5 - Правила нормирования в сфере закупок;
	6 - Требования к правовым актам по нормированию;
	7 - Требования к отдельным видам товаров, работ, услуг;
	8 - Требования к определению нормативных затрат;
	9 - Иные документы по нормированию в сфере закупок;
	10 - Правила определения требований к отдельным видам товаров, работ, услуг.
	*/
	State        zfcs_longTextType    `xml:"state"`        /* Статус. Элемент не используется в импорте */
	TermsControl bool                 `xml:"termsControl"` /* Результат контроля сроков размещения. Элемент не используется в импорте */
	ApprovedFrom `xml:"approvedFrom"` /* Орган, утвердивший правила нормирования  */
	ApproveFor   struct {
		Central     bool `xml:"central"`
		Territorial bool `xml:"territorial"`
		Treasury    bool `xml:"treasury"`
		Budgetary   bool `xml:"budgetary"`
	} `xml:"approveFor"`
	BaseDocument struct {
		Name   zfcs_longTextMinType                     `xml:"name"`
		Number zfcs_documentNumberType                  `xml:"number"`
		Date   string                                   `xml:"date"`
		Type   zfcs_regulationRulesBaseDocumentTypeEnum `xml:"type"`
	} `xml:"baseDocument"`
	Regions struct {
		Region `xml:"region"`
	} `xml:"regions"`
	AddInfo    zfcs_longTextMinType `xml:"addInfo"` /* Дополнительная информация  */
	Discussion struct {
		Term          `xml:"term"`
		PostAddress   zfcs_longTextMinType                    `xml:"postAddress"`
		EMail         string                                  `xml:"eMail"`
		ContactPerson contactPersonType                       `xml:"contactPerson"`
		Decision      zfcs_regulationRulesCouncilDecisionEnum `xml:"decision"`
		AddInfo       zfcs_longTextMinType                    `xml:"addInfo"`
	} `xml:"discussion"`
	Documents struct {
		Document `xml:"document"`
	} `xml:"documents"`
	PrintFormDocuments struct {
		PrintFormDocument `xml:"printFormDocument"`
	} `xml:"printFormDocuments"`
	Modification struct {
		Version int                  `xml:"version"`
		Info    zfcs_longTextMinType `xml:"info"`
	} `xml:"modification"`
}

type zfcs_regulationRulesInvalidType struct {
	DocPublishDate string                          `xml:"docPublishDate"` /* Дата публикации документа */
	Id             int64                           `xml:"id"`             /* Идентификатор документа ЕИС */
	RegistryNum    zfcs_standardContractNumberType `xml:"registryNum"`    /* Реестровый номер */
	RevisionNumber zfcs_revisionNumType            `xml:"revisionNumber"` /* Номер редакции */
	PublishOrg     `xml:"publishOrg"`              /* Организация, разместившая информацию */
	Type           zfcs_regulationRulesTypeEnum    `xml:"type"` /* Тип правил нормирования:
	1 - Общие правила нормирования в сфере закупок;
	2 - Общие требования к правовым актам по нормированию;
	3 - Общие требования к отдельным видам товаров, работ, услуг;
	4 - Общие требования к определению нормативных затрат;
	5 - Правила нормирования в сфере закупок;
	6 - Требования к правовым актам по нормированию;
	7 - Требования к отдельным видам товаров, работ, услуг;
	8 - Требования к определению нормативных затрат;
	9 - Иные документы по нормированию в сфере закупок;
	10 - Правила определения требований к отдельным видам товаров, работ, услуг.
	 Элемент не используется в импорте */
	State        zfcs_longTextType    `xml:"state"`        /* Статус. Элемент не используется в импорте */
	TermsControl bool                 `xml:"termsControl"` /* Результат контроля сроков размещения. Элемент не используется в импорте */
	ApprovedFrom `xml:"approvedFrom"` /* Орган, утвердивший правила нормирования. Элемент не используется в импорте  */
	ApproveFor   struct {
		Central     bool `xml:"central"`
		Territorial bool `xml:"territorial"`
		Treasury    bool `xml:"treasury"`
		Budgetary   bool `xml:"budgetary"`
	} `xml:"approveFor"`
	BaseDocument struct {
		Name   zfcs_longTextMinType                     `xml:"name"`
		Number zfcs_documentNumberType                  `xml:"number"`
		Date   string                                   `xml:"date"`
		Type   zfcs_regulationRulesBaseDocumentTypeEnum `xml:"type"`
	} `xml:"baseDocument"`
	Regions struct {
		Region `xml:"region"`
	} `xml:"regions"`
	AddInfo   zfcs_longTextMinType `xml:"addInfo"` /* Дополнительная информация. Элемент не используется в импорте */
	Documents struct {
		Document `xml:"document"`
	} `xml:"documents"`
	PrintFormDocuments struct {
		PrintFormDocument `xml:"printFormDocument"`
	} `xml:"printFormDocuments"`
	Discussion struct {
		Term        `xml:"term"`
		PostAddress zfcs_longTextMinType                    `xml:"postAddress"`
		EMail       string                                  `xml:"eMail"`
		Decision    zfcs_regulationRulesCouncilDecisionEnum `xml:"decision"`
		AddInfo     zfcs_longTextMinType                    `xml:"addInfo"`
	} `xml:"discussion"`
	InvalidityInfo struct {
		Date   string               `xml:"date"`
		Reason zfcs_longTextMinType `xml:"reason"`
	} `xml:"invalidityInfo"`
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Документы */
}

type zfcs_accountType struct {
	BankAddress     string                     `xml:"bankAddress"`     /* Адрес кредитной организации */
	BankName        string                     `xml:"bankName"`        /* Наименование кредитной организации */
	Bik             zfcs_bikType               `xml:"bik"`             /* БИК */
	CorrAccount     string                     `xml:"corrAccount"`     /* Корреспондентский счет */
	PaymentAccount  zfcs_settlementAccountType `xml:"paymentAccount"`  /* Расчетный счет */
	PersonalAccount zfcs_personalAccountType   `xml:"personalAccount"` /* Лицевой счет */
}

type zfcs_guarantee_attachmentListType struct {
	Attachment zfcs_guarantee_attachmentType `xml:"attachment"` /*  */
}

type zfcs_guarantee_attachmentType struct {
	PublishedContentId zfcs_guidType                   `xml:"publishedContentId"` /* Уникальный идентификатор контента прикрепленного документа на ЕИС  */
	FileName           string                          `xml:"fileName"`           /* Имя файла */
	FileSize           string                          `xml:"fileSize"`           /* Размер файла */
	DocDescription     string                          `xml:"docDescription"`     /* Описание прикрепляемого документа */
	RegDocNumber       zfcs_bankGuaranteeDocNumberType `xml:"regDocNumber"`       /* Реестровый номер документа */
	CryptoSigns        struct {
		Signature string `xml:"signature"`
	} `xml:"cryptoSigns"`
}

type zfcs_bankGuaranteeRefReasonType struct {
	Code int64  `xml:"code"` /* Код причины */
	Name string `xml:"name"` /* Наименование причины */
}

type zfcs_bankGuaranteeRefusalType struct {
	Id             int64                           `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType             `xml:"externalId"`     /* Внешний идентификатор документа */
	RegNumber      zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`      /* Номер реестровой записи банковской гарантии */
	DocNumber      zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`      /* Номер отказа */
	VersionNumber  int                             `xml:"versionNumber"`  /* Номер редакции отказа */
	DocPublishDate string                          `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	Bank         zfcs_bankGuaranteeOrganizationType `xml:"bank"`         /* Сведения о банке.выдавшем гарантию */
	SupplierInfo zfcs_bankGuaranteeParticipantType  `xml:"supplierInfo"` /* Информация о поставщике (подрядчике, исполнителе) */
	Guarantee    zfcs_bankGuaranteeInfoType         `xml:"guarantee"`    /* Информация о банковской гарантии */
	Href         zfcs_hrefType                      `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType                 `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType              `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	RefusalInfo  struct {
		DocDate        string `xml:"docDate"`
		DocNumber      string `xml:"docNumber"`
		DocName        string `xml:"docName"`
		RefusalReasons `xml:"refusalReasons"`
	} `xml:"refusalInfo"`
	Attachments      zfcs_attachmentListType `xml:"attachments"`      /* Прикрепленные файлы */
	ModificationInfo string                  `xml:"modificationInfo"` /* Описание изменения */
}

type zfcs_bankGuaranteeRefusalInvalidType struct {
	Id               int64                           `xml:"id"`               /* Идентификатор документа ЕИС */
	ExternalId       zfcs_externalIdType             `xml:"externalId"`       /* Внешний идентификатор документа */
	RefusalDocNumber zfcs_bankGuaranteeDocNumberType `xml:"refusalDocNumber"` /* Номер  отказа заказчика в принятии банковской гарантии */
	DocNumber        zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`        /* Номер сведений о недействительности */
	DocPublishDate   string                          `xml:"docPublishDate"`   /* Дата публикации документа
	Планируемая или фактическая */
	BankGuaranteeRefusalInfo struct {
		Bank         zfcs_bankGuaranteeOrganizationType `xml:"bank"`
		SupplierInfo zfcs_bankGuaranteeParticipantType  `xml:"supplierInfo"`
		Guarantee    zfcs_bankGuaranteeInfoType         `xml:"guarantee"`
		RefusalInfo  `xml:"refusalInfo"`
	} `xml:"bankGuaranteeRefusalInfo"`
	Href         zfcs_hrefType           `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Прикрепленные файлы */
	Reason       zfcs_longTextType       `xml:"reason"`       /* Причина недействительности */
}

type zfcs_bankGuaranteeType struct {
	Id                int64                           `xml:"id"`                /* Идентификатор документа ЕИС */
	ExternalId        zfcs_externalIdType             `xml:"externalId"`        /* Внешний идентификатор документа */
	RegNumber         zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`         /* Номер реестровой записи банковской гарантии */
	DocNumber         zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`         /* Номер документа */
	ExtendedDocNumber zfcs_bankGuaranteeDocNumberType `xml:"extendedDocNumber"` /* Последний номер документа */
	VersionNumber     int                             `xml:"versionNumber"`     /* Номер редакции */
	DocPublishDate    string                          `xml:"docPublishDate"`    /* Дата публикации документа
	Планируемая или фактическая */
	Bank         zfcs_bankGuaranteeOrganizationType `xml:"bank"`         /* Сведения о банке, выдавшем гарантию */
	PlacingOrg   zfcs_bankGuaranteeOrganizationType `xml:"placingOrg"`   /* Сведения об организации (филиале банка), разместившей гарантию */
	Supplier     zfcs_participantType               `xml:"supplier"`     /* Информация о поставщике (подрядчике, исполнителе) (устарело) */
	SupplierInfo zfcs_bankGuaranteeParticipantType  `xml:"supplierInfo"` /* Информация о поставщике (подрядчике, исполнителе) */
	Guarantee    struct {
		Customer           zfcs_bankGuaranteeOrganizationType `xml:"customer"`
		PurchaseCode       zfcs_ikzCodeType                   `xml:"purchaseCode"`
		GuaranteeDate      string                             `xml:"guaranteeDate"`
		GuaranteeNumber    zfcs_bankGuaranteeNumberType       `xml:"guaranteeNumber"`
		GuaranteeAmount    zfcs_moneyType                     `xml:"guaranteeAmount"`
		Currency           zfcs_currencyFullRef               `xml:"currency"`
		ExpireDate         string                             `xml:"expireDate"`
		EntryForceDate     string                             `xml:"entryForceDate"`
		Procedure          zfcs_longTextType                  `xml:"procedure"`
		GuaranteeAmountRUR zfcs_moneyType                     `xml:"guaranteeAmountRUR"`
		CurrencyRate       float64                            `xml:"currencyRate"`
	} `xml:"guarantee"`
	Href               zfcs_hrefType                     `xml:"href"`               /* Гиперссылка на опубликованный документ */
	PrintForm          zfcs_printFormType                `xml:"printForm"`          /* Печатная форма документа */
	AgreementDocuments zfcs_guarantee_attachmentListType `xml:"agreementDocuments"` /* Копия заключенного договора банковской гарантии */
	Modification       struct {
		ModificationDate      string                            `xml:"modificationDate"`
		Reason                zfcs_bankGaranteeModificationType `xml:"reason"`
		Info                  string                            `xml:"info"`
		ModificationDocuments zfcs_guarantee_attachmentListType `xml:"modificationDocuments"`
	} `xml:"modification"`
}

type zfcs_bankGuaranteeInvalidType struct {
	Id             int64                           `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType             `xml:"externalId"`     /* Внешний идентификатор документа */
	RegNumber      zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`      /* Номер реестровой записи банковской гарантии */
	DocNumber      zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`      /* Номер документа */
	DocPublishDate string                          `xml:"docPublishDate"` /* Дата публикации документа */
	Reason         zfcs_longTextType               `xml:"reason"`         /* Причина недействительности */
	GuaranteeInfo  struct {
		Bank         zfcs_bankGuaranteeOrganizationType `xml:"bank"`
		PlacingOrg   zfcs_bankGuaranteeOrganizationType `xml:"placingOrg"`
		SupplierInfo zfcs_bankGuaranteeParticipantType  `xml:"supplierInfo"`
		Guarantee    zfcs_bankGuaranteeInfoType         `xml:"guarantee"`
	} `xml:"guaranteeInfo"`
	Href      zfcs_hrefType           `xml:"href"`      /* Гиперссылка на опубликованный документ */
	PrintForm zfcs_printFormType      `xml:"printForm"` /* Печатная форма документа */
	Documents zfcs_attachmentListType `xml:"documents"` /* Документы */
}

type zfcs_bankGuaranteeTerminationType struct {
	Id             int64                           `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType             `xml:"externalId"`     /* Внешний идентификатор документа */
	RegNumber      zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`      /* Номер реестровой записи банковской гарантии */
	DocNumber      zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`      /* Номер информации о прекращении */
	VersionNumber  int                             `xml:"versionNumber"`  /* Номер редакции информации о прекращении */
	DocPublishDate string                          `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	Bank                     zfcs_bankGuaranteeOrganizationType                     `xml:"bank"`                     /* Сведения о банке.выдавшем гарантию */
	SupplierInfo             zfcs_bankGuaranteeParticipantType                      `xml:"supplierInfo"`             /* Информация о поставщике (подрядчике, исполнителе) */
	Guarantee                zfcs_bankGuaranteeInfoType                             `xml:"guarantee"`                /* Информация о банковской гарантии */
	RegNum                   zfcs_contract_regNumType                               `xml:"regNum"`                   /* Номер реестровой записи контракта */
	BankGuaranteeTermination zfcs_contractProcedure2015BankGuaranteeTerminationType `xml:"bankGuaranteeTermination"` /* Информация о прекращении обязательств поставщика, обеспеченных банковской гарантией */
	Href                     zfcs_hrefType                                          `xml:"href"`                     /* Гиперссылка на опубликованный документ */
	PrintForm                zfcs_printFormType                                     `xml:"printForm"`                /* Печатная форма документа */
	ExtPrintForm             zfcs_extPrintFormType                                  `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	Attachments              zfcs_attachmentListType                                `xml:"attachments"`              /* Прикрепленные файлы */
	ModificationInfo         string                                                 `xml:"modificationInfo"`         /* Описание изменения */
}

type zfcs_bankGuaranteeTerminationInvalidType struct {
	Id                   int64                           `xml:"id"`                   /* Идентификатор документа ЕИС */
	ExternalId           zfcs_externalIdType             `xml:"externalId"`           /* Внешний идентификатор документа */
	TerminationDocNumber zfcs_bankGuaranteeDocNumberType `xml:"terminationDocNumber"` /* Номер  информации о прекращении обязательств поставщика по банковской гарантии */
	DocNumber            zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`            /* Номер сведений о недействительности */
	DocPublishDate       string                          `xml:"docPublishDate"`       /* Дата публикации документа
	Планируемая или фактическая */
	BankGuaranteeTerminationInfo struct {
		Bank                     zfcs_bankGuaranteeOrganizationType `xml:"bank"`
		SupplierInfo             zfcs_bankGuaranteeParticipantType  `xml:"supplierInfo"`
		Guarantee                zfcs_bankGuaranteeInfoType         `xml:"guarantee"`
		BankGuaranteeTermination `xml:"bankGuaranteeTermination"`
	} `xml:"bankGuaranteeTerminationInfo"`
	Href         zfcs_hrefType           `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Прикрепленные файлы */
	Reason       zfcs_longTextType       `xml:"reason"`       /* Причина недействительности */
}

type zfcs_bankGuaranteeReturnType struct {
	Id             int64                           `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType             `xml:"externalId"`     /* Внешний идентификатор документа */
	RegNumber      zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`      /* Номер реестровой записи банковской гарантии */
	DocNumber      zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`      /* Номер информации о возвращении */
	VersionNumber  int                             `xml:"versionNumber"`  /* Номер редакции информации о возвращении */
	DocPublishDate string                          `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	Bank             zfcs_bankGuaranteeOrganizationType       `xml:"bank"`             /* Сведения о банке.выдавшем гарантию */
	SupplierInfo     zfcs_bankGuaranteeParticipantType        `xml:"supplierInfo"`     /* Информация о поставщике (подрядчике, исполнителе) */
	Guarantee        zfcs_bankGuaranteeInfoType               `xml:"guarantee"`        /* Информация о банковской гарантии */
	RegNum           zfcs_contract_regNumType                 `xml:"regNum"`           /* Номер реестровой записи контракта */
	GuaranteeReturns zfcs_contract2015BankGuaranteeReturnType `xml:"guaranteeReturns"` /* Информация о возвращении банковской гарантии или уведомление об освобождении от обязательств по банковской гарантии */
	Href             zfcs_hrefType                            `xml:"href"`             /* Гиперссылка на опубликованный документ */
	PrintForm        zfcs_printFormType                       `xml:"printForm"`        /* Печатная форма документа */
	ExtPrintForm     zfcs_extPrintFormType                    `xml:"extPrintForm"`     /* Электронный документ, полученный из внешней системы */
	Attachments      zfcs_attachmentListType                  `xml:"attachments"`      /* Прикрепленные файлы */
	ModificationInfo string                                   `xml:"modificationInfo"` /* Описание изменения */
}

type zfcs_bankGuaranteeReturnInvalidType struct {
	Id              int64                           `xml:"id"`              /* Идентификатор документа ЕИС */
	ExternalId      zfcs_externalIdType             `xml:"externalId"`      /* Внешний идентификатор документа */
	ReturnDocNumber zfcs_bankGuaranteeDocNumberType `xml:"returnDocNumber"` /* Номер  информации о возвращении банковской гарантии или об освобождении от обязательств по банковской гарантии */
	DocNumber       zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`       /* Номер сведений о недействительности */
	DocPublishDate  string                          `xml:"docPublishDate"`  /* Дата публикации документа
	Планируемая или фактическая */
	BankGuaranteeReturnInfo struct {
		Bank             zfcs_bankGuaranteeOrganizationType       `xml:"bank"`
		SupplierInfo     zfcs_bankGuaranteeParticipantType        `xml:"supplierInfo"`
		Guarantee        zfcs_bankGuaranteeInfoType               `xml:"guarantee"`
		GuaranteeReturns zfcs_contract2015BankGuaranteeReturnType `xml:"guaranteeReturns"`
	} `xml:"bankGuaranteeReturnInfo"`
	Href         zfcs_hrefType           `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Прикрепленные файлы */
	Reason       zfcs_longTextType       `xml:"reason"`       /* Причина недействительности */
}

type zfcs_bankGuaranteeOrganizationType struct {
	RegNum           zfcs_string              `xml:"regNum"`           /* Код по СПЗ */
	ConsRegistryNum  zfcs_consRegistryNumType `xml:"consRegistryNum"`  /* Код по Сводному Реестру */
	FullName         zfcs_longTextType        `xml:"fullName"`         /* Полное наименование */
	ShortName        zfcs_longTextType        `xml:"shortName"`        /* Сокращенное наименование */
	PostAddress      zfcs_longTextType        `xml:"postAddress"`      /* Почтовый адрес организации */
	FactAddress      zfcs_longTextType        `xml:"factAddress"`      /* Адрес местонахождения организации */
	INN              zfcs_innOrganizationType `xml:"INN"`              /* ИНН организации */
	KPP              zfcs_kppType             `xml:"KPP"`              /* КПП организации */
	Location         zfcs_longTextType        `xml:"location"`         /* Адрес места нахождения банка */
	LegalForm        okopfType                `xml:"legalForm"`        /* Организационно-правовая форма организации в ОКОПФ */
	SubjectRF        zfcs_subjectRFRef        `xml:"subjectRF"`        /* Субъект РФ */
	OKTMO            zfcs_OKTMORef            `xml:"OKTMO"`            /* ОКТМО */
	RegistrationDate string                   `xml:"registrationDate"` /* Дата постановки на учет в налоговом органе */
	IKU              zfcs_ikuType             `xml:"IKU"`              /* ИКУ организации */
}

type zfcs_bankGuaranteeParticipantType struct {
}

type zfcs_bankGuaranteeInfoType struct {
	Customer             zfcs_bankGuaranteeOrganizationType `xml:"customer"`             /* Сведения о заказчике */
	GuaranteeDate        string                             `xml:"guaranteeDate"`        /* Дата банковской гарантии */
	GuaranteeGrantDate   string                             `xml:"guaranteeGrantDate"`   /* Дата выдачи банковской гарантии */
	GuaranteePublishDate string                             `xml:"guaranteePublishDate"` /* Дата публикации банковской гарантии */
	GuaranteeNumber      zfcs_bankGuaranteeNumberType       `xml:"guaranteeNumber"`      /* Номер банковской гарантии (Не используется) */
	ExpireDate           string                             `xml:"expireDate"`           /* Дата окончания срока действия банковской гарантии */
	GuaranteeAmount      zfcs_moneyType                     `xml:"guaranteeAmount"`      /* Сумма */
	Currency             zfcs_currencyFullRef               `xml:"currency"`             /* Валюта */
	EntryForceDate       string                             `xml:"entryForceDate"`       /* Дата вступления в силу банковской гарантии */
	Procedure            zfcs_longTextType                  `xml:"procedure"`            /* Порядок вступления в силу банковской гарантии */
	GuaranteeAmountRUR   zfcs_moneyType                     `xml:"guaranteeAmountRUR"`   /* Сумма в рублях */
	CurrencyRate         float64                            `xml:"currencyRate"`         /* Курс валюты по отношению к рублю */
}

type zfcs_contract2015SupplierType struct {
}

type zfcs_contract2015SingleCustomerType struct {
	Reason struct {
		Code zfcs_nsiSingleCustomerReasonCodeType `xml:"code"`
		Name zfcs_nsiSingleCustomerReasonNameType `xml:"name"`
	} `xml:"reason"`
	Document struct {
		Code zfcs_nsiSingleCustomerReasonDocCodeType `xml:"code"`
		Name zfcs_nsiSingleCustomerReasonDocNameType `xml:"name"`
	} `xml:"document"`
	ReportBase  zfcs_longTextMinType             `xml:"reportBase"`  /* Реквизиты отчета */
	ReportCode  string                           `xml:"reportCode"`  /* Код отчета (для печатной формы) */
	Attachments zfcs_contract_attachmentListType `xml:"attachments"` /* Документы, подтверждающие основание заключения контракта */
}

type zfcs_confirmationType struct {
	LoadId string                      `xml:"loadId"` /* Идентификатор загруженных данных */
	Result zfcs_confirmationResultType `xml:"result"` /* Результат обработки данных

	success - обработано успешно;
	failure - ошибки при обработке;
	processing - обрабатывается. */
	Violations struct {
		Violation zfcs_violationType `xml:"violation"`
	} `xml:"violations"`
	LoadUrl string `xml:"loadUrl"` /* Гиперссылка на форму отображения загруженных данных */
	RefId   string `xml:"refId"`   /* Идентификатор обработанного пакета данных */
}

type zfcs_signIncomingConfirmationType struct {
	RefPacketUid zfcs_guidType `xml:"refPacketUid"` /* Идентификатор обработанного пакета данных */
}

type zfcs_signIncomingReqConfirmationType struct {
	RefPacketUid zfcs_guidType `xml:"refPacketUid"` /* Идентификатор переданного на ЕИС пакета данных */
}

type zfcs_printFormType struct {
	Url       string `xml:"url"`       /* Ссылка для скачивания печатной формы */
	Signature string `xml:"signature"` /* Электронная подпись печатной формы */
}

type zfcs_extPrintFormType struct {
	Signature                `xml:"signature"`                /* Электронная подпись электронного документа */
	FileType                 zfcs_printFormFileType           `xml:"fileType"` /* Тип файла электронного документа */
	ControlPersonalSignature `xml:"controlPersonalSignature"` /* Электронная подпись электронного документа лицом, уполномоченным на проведение контроля в соответствии с ч.5 ст.99 закона №44-ФЗ */
}

type zfcs_docType struct {
	Code string `xml:"code"` /* Кодовое наименование типа */
	Name string `xml:"name"` /* Наименование типа документа */
}

type zfcs_applicantType struct {
	ApplicantType zfcs_complaintApplicantType `xml:"applicantType"` /* Тип заявителя:
	P - Физическое лицо;
	U - Юридическое лицо;
	I - Индивидуальный предприниматель */
	OrganizationName zfcs_longTextType `xml:"organizationName"` /* Наименование юр. лица/Фамилия Имя Отчество для физ.лица */
	FactualAddress   string            `xml:"factualAddress"`   /* Место нахождения */
	PostAddress      string            `xml:"postAddress"`      /* Почтовый адрес */
	ContactEMail     string            `xml:"contactEMail"`     /* Адрес электронной почты */
	ContactPhone     zfcs_string       `xml:"contactPhone"`     /* Номер контактного телефона */
	ContactFax       zfcs_string       `xml:"contactFax"`       /* Факс  */
}

type zfcs_checkPlanType struct {
	CommonInfo struct {
		CheckPlanNumber zfcs_checkPlanNumberType `xml:"checkPlanNumber"`
		VersionNumber   int                      `xml:"versionNumber"`
		CreateDate      string                   `xml:"createDate"`
		ConfirmDate     string                   `xml:"confirmDate"`
		PublishDate     string                   `xml:"publishDate"`
		Owner           zfcs_organizationRef     `xml:"owner"`
	} `xml:"commonInfo"`
	StartStage zfcs_stageType `xml:"startStage"` /* Начало периода планирования */
	EndStage   zfcs_stageType `xml:"endStage"`   /* Окончание периода планирования */
	CheckList  struct {
		CheckInfo `xml:"checkInfo"`
	} `xml:"checkList"`
	PrintForm    zfcs_printFormType    `xml:"printForm"`    /* Печатная форма плана проверок */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_checkResultActType struct {
	ActNumber   string                  `xml:"actNumber"`   /* Номер акта */
	ActDate     string                  `xml:"actDate"`     /* Дата принятия акта */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
}

type zfcs_checkResultCancelType struct {
	CommonInfo struct {
		CheckResultNumber zfcs_checkResultNumberType `xml:"checkResultNumber"`
		CreateDate        string                     `xml:"createDate"`
		PublishDate       string                     `xml:"publishDate"`
		Owner             zfcs_organizationRef       `xml:"owner"`
	} `xml:"commonInfo"`
	CancelType zfcs_checkResultCancel_cancelType `xml:"cancelType"` /* Тип отмены результата контроля
	CO_DECISION - по решению контролирующего органа;
	JUDGE_DECISION - по решению судебного органа. */
	Info         zfcs_longTextType       `xml:"info"`         /* Комментарий */
	DocumentName string                  `xml:"documentName"` /* Наименование документа */
	DocumentDate string                  `xml:"documentDate"` /* Дата документа */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма информации об отмене */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
}

type zfcs_checkResultDecisionType struct {
	DecisionNumber string                  `xml:"decisionNumber"` /* Номер */
	DecisionDate   string                  `xml:"decisionDate"`   /* Дата принятия  */
	Attachments    zfcs_attachmentListType `xml:"attachments"`    /* Информация о прикрепленных документах */
}

type zfcs_checkResultPrescriptionType struct {
	PrescriptionNumber zfcs_prescriptionNumberType `xml:"prescriptionNumber"` /* Номер предписания */
	PrescriptionDate   string                      `xml:"prescriptionDate"`   /* Дата принятия предписания */
	Attachments        zfcs_attachmentListType     `xml:"attachments"`        /* Информация о прикрепленных документах */
}

type zfcs_checkResultType struct {
	CommonInfo struct {
		CheckResultNumber zfcs_checkResultNumberType  `xml:"checkResultNumber"`
		VersionNumber     int                         `xml:"versionNumber"`
		CreateDate        string                      `xml:"createDate"`
		PublishDate       string                      `xml:"publishDate"`
		Owner             zfcs_organizationRef        `xml:"owner"`
		Result            zfcs_checkResult_resultType `xml:"result"`
	} `xml:"commonInfo"`
	StartDate    string                `xml:"startDate"`    /* Дата начала проверки */
	EndDate      string                `xml:"endDate"`      /* Дата окончания проверки */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_checkSubjectPlanType struct {
}

type zfcs_complaintCancelType struct {
	ComplaintNumber zfcs_complaintNumberType `xml:"complaintNumber"` /* Номер жалобы */
	IsGroupItem     bool                     `xml:"isGroupItem"`     /* Признак, входит ли отзываемая жалоба в группу жалоб */
	RegDate         string                   `xml:"regDate"`         /* Дата регистрации отзыва */
	RegistrationKO  zfcs_organizationRef     `xml:"registrationKO"`  /* Орган, осуществлявший регистрацию отзыва жалобы */
	Name            zfcs_longTextType        `xml:"name"`            /* Наименование участника контрактной системы в сфере закупок, отозвавшего жалобу */
	RegType         zfcs_complaint_regType   `xml:"regType"`         /* Тип подачи жалобы
	M-ручное заведение жалобы
	E - электронная подача
	I - полученная по интеграции */
	Text         zfcs_longTextType       `xml:"text"`         /* Основание отзыва жалобы */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма отзыва жалобы */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
}

type zfcs_complaintCommonInfoType struct {
	ComplaintNumber  zfcs_complaintNumberType `xml:"complaintNumber"`  /* Номер жалобы */
	VersionNumber    int                      `xml:"versionNumber"`    /* Номер редакции */
	PlanDecisionDate string                   `xml:"planDecisionDate"` /* Дата и время рассмотрения жалобы */
	DecisionPlace    zfcs_longTextType        `xml:"decisionPlace"`    /* Место рассмотрения жалобы */
	RegistrationKO   zfcs_organizationRef     `xml:"registrationKO"`   /* Орган, осуществлявший регистрацию жалобы */
	ConsiderationKO  zfcs_organizationRef     `xml:"considerationKO"`  /* Орган, осуществлявший рассмотрение жалобы */
	RegType          zfcs_complaint_regType   `xml:"regType"`          /* Тип подачи жалобы
	M-ручное заведение жалобы
	E - электронная подача
	I - полученная по интеграции */
}

type zfcs_complaintGroupType struct {
	CommonInfo zfcs_complaintCommonInfoType `xml:"commonInfo"` /* Общая информация о группой жалобе */
	Indicted   struct {
		EP_failure zfcs_organizationRef `xml:"EP_failure"`
	} `xml:"indicted"`
	Text                zfcs_longTextType `xml:"text"` /* Содержание жалобы (обжалуемые действия) */
	ComplaintGroupItems struct {
		ItemComplaintNumber zfcs_complaintNumberType `xml:"itemComplaintNumber"`
		RegDate             string                   `xml:"regDate"`
		Applicant           zfcs_applicantType       `xml:"applicant"`
		ReturnInfo          `xml:"returnInfo"`
		Attachments         zfcs_attachmentListType `xml:"attachments"`
	} `xml:"complaintGroupItems"`
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма жалобы */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
	ReturnInfo   struct {
		Base        string                  `xml:"base"`
		Attachments zfcs_attachmentListType `xml:"attachments"`
	} `xml:"returnInfo"`
}

type zfcs_complaintObjectType struct {
}

type zfcs_complaintOrderType struct {
	NotificationNumber zfcs_string94 `xml:"notificationNumber"` /* Номер извещения о проведении */
	Lots               struct {
		LotNumber int               `xml:"lotNumber"`
		Info      zfcs_longTextType `xml:"info"`
	} `xml:"lots"`
}

type zfcs_complaintPurchaseType struct {
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер извещения о проведении */
	Lots           struct {
		LotNumber int               `xml:"lotNumber"`
		Info      zfcs_longTextType `xml:"info"`
	} `xml:"lots"`
}

type zfcs_complaintSubjectType struct {
}

type zfcs_complaintProjectSubjectType struct {
}

type zfcs_complaintType struct {
	CommonInfo `xml:"commonInfo"`        /* Общая информация о жалобе */
	Indicted   zfcs_complaintSubjectType `xml:"indicted"`  /* На кого подана жалоба, субъекты */
	Applicant  zfcs_applicantType        `xml:"applicant"` /* Заявитель */
	Object     zfcs_complaintObjectType  `xml:"object"`    /* Предмет жалобы, допустимые значения
	Закупка/ заказ
	План закупки
	План-график
	*/
	Text         zfcs_longTextType            `xml:"text"`         /* Содержание жалобы (обжалуемые действия) */
	PrintForm    zfcs_printFormType           `xml:"printForm"`    /* Печатная форма жалобы */
	ExtPrintForm zfcs_extPrintFormType        `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType      `xml:"attachments"`  /* Информация о прикрепленных документах */
	ReturnInfo   zfcs_complaintReturnInfoType `xml:"returnInfo"`   /* Сведения о возврате жалобы */
}

type zfcs_complaintProjectType struct {
	Id         int64               `xml:"id"`         /* Идентификатор документа ЕИС */
	ExternalId zfcs_externalIdType `xml:"externalId"` /* Внешний идентификатор документа */
	CommonInfo struct {
		RegDate          string                   `xml:"regDate"`
		ComplaintNumber  zfcs_complaintNumberType `xml:"complaintNumber"`
		VersionNumber    int                      `xml:"versionNumber"`
		PlanDecisionDate string                   `xml:"planDecisionDate"`
		DecisionPlace    zfcs_longTextType        `xml:"decisionPlace"`
		RegistrationKO   zfcs_organizationRef     `xml:"registrationKO"`
		ConsiderationKO  zfcs_organizationRef     `xml:"considerationKO"`
		RegType          zfcs_complaint_regType   `xml:"regType"`
	} `xml:"commonInfo"`
	Indicted  zfcs_complaintProjectSubjectType `xml:"indicted"`  /* На кого подана жалоба, субъекты */
	Applicant zfcs_applicantType               `xml:"applicant"` /* Заявитель */
	Object    zfcs_complaintObjectType         `xml:"object"`    /* Предмет жалобы, допустимые значения
	Закупка/ заказ;
	План закупки;
	План-график.  */
	Text         zfcs_longTextType            `xml:"text"`         /* Содержание жалобы (обжалуемые действия) */
	PrintForm    zfcs_printFormType           `xml:"printForm"`    /* Печатная форма жалобы */
	ExtPrintForm zfcs_extPrintFormType        `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType      `xml:"attachments"`  /* Информация о прикрепленных документах */
	ReturnInfo   zfcs_complaintReturnInfoType `xml:"returnInfo"`   /* Сведения о возврате жалобы */
}

type zfcs_electronicComplaintType struct {
	Id         int64               `xml:"id"`         /* Идентификатор документа ЕИС */
	ExternalId zfcs_externalIdType `xml:"externalId"` /* Внешний идентификатор документа */
	CommonInfo struct {
		RegDate          string                   `xml:"regDate"`
		ComplaintNumber  zfcs_complaintNumberType `xml:"complaintNumber"`
		VersionNumber    int                      `xml:"versionNumber"`
		PlanDecisionDate string                   `xml:"planDecisionDate"`
		DecisionPlace    zfcs_longTextType        `xml:"decisionPlace"`
		RegistrationKO   zfcs_organizationRef     `xml:"registrationKO"`
		ConsiderationKO  zfcs_organizationRef     `xml:"considerationKO"`
		RegType          zfcs_complaint_regType   `xml:"regType"`
	} `xml:"commonInfo"`
	Indicted  zfcs_complaintProjectSubjectType `xml:"indicted"`  /* На кого подана жалоба, субъекты */
	Applicant zfcs_applicantType               `xml:"applicant"` /* Заявитель */
	Object    zfcs_complaintObjectType         `xml:"object"`    /* Предмет жалобы, допустимые значения
	Закупка/ заказ
	План закупки
	План-график
	*/
	Text         zfcs_longTextType            `xml:"text"`         /* Содержание жалобы (обжалуемые действия) */
	PrintForm    zfcs_printFormType           `xml:"printForm"`    /* Печатная форма жалобы */
	ExtPrintForm zfcs_extPrintFormType        `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType      `xml:"attachments"`  /* Информация о прикрепленных документах */
	ReturnInfo   zfcs_complaintReturnInfoType `xml:"returnInfo"`   /* Сведения о возврате жалобы */
}

type zfcs_complaintReturnInfoType struct {
	Base        string                  `xml:"base"`        /* Основание возврата жалобы */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
}

type zfcs_complaintCancelInfoType struct {
	Name        zfcs_longTextType       `xml:"name"`        /* Наименование участника контрактной системы в сфере закупок, отозвавшего жалобу */
	Text        zfcs_longTextType       `xml:"text"`        /* Основание отзыва жалобы */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
}

type zfcs_periodType struct {
	Start string `xml:"start"` /* Проверяемый период  "С" */
	End   string `xml:"end"`   /* Проверяемый период "По" */
}

type zfcs_unplannedCheckCancelType struct {
	CheckNumber  zfcs_unplannedCheckNumberType `xml:"checkNumber"`  /* Номер внеплановой проверки */
	PublishDate  string                        `xml:"publishDate"`  /* Дата публикации отмены */
	Text         zfcs_longTextType             `xml:"text"`         /* Причина отмены проверки */
	PrintForm    zfcs_printFormType            `xml:"printForm"`    /* Печатная форма отмены внеплановой проверки */
	ExtPrintForm zfcs_extPrintFormType         `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Инфрмация о прикрепленных документах */
}

type zfcs_unplannedCheckObjectType struct {
}

type zfcs_unplannedCheckSubjectPlanType struct {
}

type zfcs_unplannedCheckType struct {
	CommonInfo struct {
		CheckNumber   zfcs_unplannedCheckNumberType `xml:"checkNumber"`
		VersionNumber int                           `xml:"versionNumber"`
		CreateDate    string                        `xml:"createDate"`
		PublishDate   string                        `xml:"publishDate"`
	} `xml:"commonInfo"`
	Period          zfcs_periodType                    `xml:"period"`          /* Проверяемый период  */
	Inspector       zfcs_organizationRef               `xml:"inspector"`       /* Орган, осуществляющий проведение проверки */
	InspectionDate  string                             `xml:"inspectionDate"`  /* Дата и время заседания */
	InspectionPlace zfcs_longTextType                  `xml:"inspectionPlace"` /* Место заседания инспекции */
	CheckedSubject  zfcs_unplannedCheckSubjectPlanType `xml:"checkedSubject"`  /* Субъект внеплановой проверки */
	Base            `xml:"base"`                       /* Основание проведения проверки */
	CheckedObject   struct {
		CheckedOrder    `xml:"checkedOrder"`
		ObjectOtherInfo zfcs_longTextType `xml:"objectOtherInfo"`
		Info            zfcs_longTextType `xml:"info"`
	} `xml:"checkedObject"`
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма внеплановой проверки */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
}

type zfcs_eventPlanType struct {
	CommonInfo struct {
		EventPlanNumber zfcs_eventPlanNumberType `xml:"eventPlanNumber"`
		VersionNumber   int                      `xml:"versionNumber"`
		CreateDate      string                   `xml:"createDate"`
		ConfirmDate     string                   `xml:"confirmDate"`
		PublishDate     string                   `xml:"publishDate"`
		Owner           zfcs_organizationRef     `xml:"owner"`
	} `xml:"commonInfo"`
	StartStage zfcs_stageType `xml:"startStage"` /* Начало периода планирования */
	EndStage   zfcs_stageType `xml:"endStage"`   /* Окончание периода планирования */
	EventList  struct {
		EventInfo `xml:"eventInfo"`
	} `xml:"eventList"`
	PrintForm    zfcs_printFormType    `xml:"printForm"`    /* Печатная форма плана мероприятий */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_unplannedEventType struct {
	CommonInfo struct {
		EventNumber   zfcs_unplannedEventNumberType `xml:"eventNumber"`
		VersionNumber int                           `xml:"versionNumber"`
		CreateDate    string                        `xml:"createDate"`
		PublishDate   string                        `xml:"publishDate"`
	} `xml:"commonInfo"`
	UnplannedEventType `xml:"unplannedEventType"` /* Вид внепланового контрольного мероприятия */
	Period             zfcs_periodType            `xml:"period"`         /* Проверяемый период  */
	Inspector          zfcs_organizationRef       `xml:"inspector"`      /* Орган, осуществляющий проведение мероприятия */
	CheckedSubject     zfcs_eventSubjectType      `xml:"checkedSubject"` /* Субъект контрольного мероприятия */
	Base               `xml:"base"`               /* Основание проведения мероприятия */
	CheckedObject      struct {
		Info zfcs_longTextType `xml:"info"`
	} `xml:"checkedObject"`
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма контрольного мероприятия */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
}

type zfcs_unplannedEventCancelType struct {
	EventNumber  zfcs_unplannedEventNumberType `xml:"eventNumber"`  /* Номер внепланового контрольного мероприятия */
	PublishDate  string                        `xml:"publishDate"`  /* Дата и время публикации отмены */
	Text         zfcs_longTextType             `xml:"text"`         /* Причина отмены мероприятия */
	PrintForm    zfcs_printFormType            `xml:"printForm"`    /* Печатная форма отмены мероприятия */
	ExtPrintForm zfcs_extPrintFormType         `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Инфрмация о прикрепленных документах */
}

type zfcs_eventResultType struct {
	CommonInfo struct {
		EventResultNumber zfcs_eventResultNumberType  `xml:"eventResultNumber"`
		VersionNumber     int                         `xml:"versionNumber"`
		CreateDate        string                      `xml:"createDate"`
		PublishDate       string                      `xml:"publishDate"`
		Owner             zfcs_organizationRef        `xml:"owner"`
		Result            zfcs_eventResult_resultType `xml:"result"`
	} `xml:"commonInfo"`
	StartDate    string                `xml:"startDate"`    /* Дата начала мероприятия */
	EndDate      string                `xml:"endDate"`      /* Дата окончания мероприятия */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_eventResultCancelType struct {
	CommonInfo struct {
		EventResultNumber zfcs_eventResultNumberType `xml:"eventResultNumber"`
		CreateDate        string                     `xml:"createDate"`
		PublishDate       string                     `xml:"publishDate"`
		Owner             zfcs_organizationRef       `xml:"owner"`
	} `xml:"commonInfo"`
	CancelType zfcs_eventResultCancel_cancelType `xml:"cancelType"` /* Тип отмены результата контроля
	CO_DECISION - по решению контролирующего органа;
	JUDGE_DECISION - по решению судебного органа. */
	Info         zfcs_longTextType       `xml:"info"`         /* Комментарий */
	DocumentName string                  `xml:"documentName"` /* Наименование документа */
	DocumentDate string                  `xml:"documentDate"` /* Дата документа */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма информации об отмене */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
}

type zfcs_eventSubjectType struct {
}

type zfcs_eventObjectType struct {
}

type zfcs_eventResultActType struct {
	ActNumber   string                  `xml:"actNumber"`   /* Номер акта */
	ActDate     string                  `xml:"actDate"`     /* Дата принятия акта */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
}

type zfcs_eventResultDecisionType struct {
	DecisionNumber string                  `xml:"decisionNumber"` /* Номер */
	DecisionDate   string                  `xml:"decisionDate"`   /* Дата принятия  */
	Attachments    zfcs_attachmentListType `xml:"attachments"`    /* Информация о прикрепленных документах */
}

type zfcs_eventResultPrescriptionType struct {
	PrescriptionNumber zfcs_prescriptionNumberType `xml:"prescriptionNumber"` /* Номер предписания */
	PrescriptionDate   string                      `xml:"prescriptionDate"`   /* Дата принятия предписания */
	Attachments        zfcs_attachmentListType     `xml:"attachments"`        /* Информация о прикрепленных документах */
}

type zfcs_contract_attachmentListType struct {
	Attachment zfcs_contract_attachmentType `xml:"attachment"` /*  */
}

type zfcs_contract_attachmentType struct {
	PublishedContentId zfcs_guidType               `xml:"publishedContentId"` /* Уникальный идентификатор контента прикрепленного документа на ЕИС  */
	FileName           string                      `xml:"fileName"`           /* Имя файла */
	DocDescription     string                      `xml:"docDescription"`     /* Описание прикрепляемого документа */
	DocRegNumber       zfcs_contract_docRegNumType `xml:"docRegNumber"`       /* Реестровый номер документа */
	CryptoSigns        struct {
		Signature string `xml:"signature"`
	} `xml:"cryptoSigns"`
}

type zfcs_contract_printFormType struct {
	Url          string                      `xml:"url"`          /* Ссылка для скачивания печатной формы */
	DocRegNumber zfcs_contract_docRegNumType `xml:"docRegNumber"` /* Реестровый номер документа */
	Signature    `xml:"signature"`           /* Электронная подпись печатной формы */
}

type zfcs_contract_OKEIType struct {
	Code         string `xml:"code"`         /* Код */
	NationalCode string `xml:"nationalCode"` /* Наименование */
}

type zfcs_contract2015_documentInfo struct {
	DocumentName string `xml:"documentName"` /* Наименование документа */
	DocumentNum  string `xml:"documentNum"`  /* Номер документа */
	DocumentDate string `xml:"documentDate"` /* Дата документа */
}

type zfcs_contract2015_payDocInfo struct {
	DocumentName string         `xml:"documentName"` /* Наименование */
	DocumentNum  string         `xml:"documentNum"`  /* Номер документа */
	DocumentDate string         `xml:"documentDate"` /* Дата документа */
	Amount       zfcs_moneyType `xml:"amount"`       /* Сумма в валюте */
	AmountRUR    zfcs_moneyType `xml:"amountRUR"`    /* Сумма в рублевом эквиваленте */
}

type zfcs_contract2015_DocDictRef struct {
	Code         string `xml:"code"`         /* Код */
	Name         string `xml:"name"`         /* Наименование документа */
	DocumentDate string `xml:"documentDate"` /* Дата документа */
	DocumentNum  string `xml:"documentNum"`  /* Номер документа */
}

type zfcs_contractCancelType struct {
	RegNum               zfcs_contract_regNumType        `xml:"regNum"`               /* Номер реестровой записи государственного или муниципального контракта */
	CancelDate           string                          `xml:"cancelDate"`           /* Дата аннулирования контракта */
	DocumentBase         string                          `xml:"documentBase"`         /* Реквизиты документа, подтверждающего аннулирование контракта */
	CurrentContractStage zfcs_contract_contractStageType `xml:"currentContractStage"` /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
}

type zfcs_contractCancel2015Type struct {
	RegNum               zfcs_contract_regNumType        `xml:"regNum"`               /* Номер реестровой записи государственного или муниципального контракта */
	CancelDate           string                          `xml:"cancelDate"`           /* Дата аннулирования контракта */
	DocumentBase         zfcs_longTextMinType            `xml:"documentBase"`         /* Реквизиты документа, подтверждающего аннулирование контракта */
	CurrentContractStage zfcs_contract_contractStageType `xml:"currentContractStage"` /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
}

type zfcs_contractMultiType struct {
	ContractCount uint64 `xml:"contractCount"` /* Количество контрактов */
}

type zfcs_contractProcedureCancelType struct {
	CancelledProcedureId int64                           `xml:"cancelledProcedureId"` /* Идентификатор отменяемой информации об исполнении (расторжении) контракта */
	RegNum               zfcs_contract_regNumType        `xml:"regNum"`               /* Номер реестровой записи государственного или муниципального контракта */
	CancelDate           string                          `xml:"cancelDate"`           /* Дата отмены информации об исполнении (расторжении) контракта */
	Reason               string                          `xml:"reason"`               /* Основание отмены информации об исполнении (расторжении) контракта */
	CurrentContractStage zfcs_contract_contractStageType `xml:"currentContractStage"` /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
}

type zfcs_contractProcedureCancel2015Type struct {
	CancelledProcedureId int64                           `xml:"cancelledProcedureId"` /* Идентификатор отменяемой информации об исполнении (расторжении) контракта */
	RegNum               zfcs_contract_regNumType        `xml:"regNum"`               /* Номер реестровой записи государственного или муниципального контракта */
	CancelDate           string                          `xml:"cancelDate"`           /* Дата отмены информации об исполнении (расторжении) контракта */
	Reason               zfcs_longTextMinType            `xml:"reason"`               /* Основание отмены информации об исполнении (расторжении) контракта */
	ExtPrintForm         zfcs_extPrintFormType           `xml:"extPrintForm"`         /* Электронный документ, полученный из внешней системы */
	CurrentContractStage zfcs_contract_contractStageType `xml:"currentContractStage"` /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
}

type zfcs_contractProcedureType struct {
	Id            int64                    `xml:"id"`            /* Идентификатор документа ЕИС */
	ExternalId    zfcs_externalIdType      `xml:"externalId"`    /* Внешний идентификатор документа */
	RegNum        zfcs_contract_regNumType `xml:"regNum"`        /* Номер реестровой записи государственного или муниципального контракта */
	PublishDate   string                   `xml:"publishDate"`   /* Дата публикации */
	VersionNumber int                      `xml:"versionNumber"` /* Номер редакции сведений */
	Executions    struct {
		Stage               zfcs_stageType `xml:"stage"`
		OrdinalNumber       int            `xml:"ordinalNumber"`
		FinalStageExecution bool           `xml:"finalStageExecution"`
		Execution           `xml:"execution"`
	} `xml:"executions"`
	Terminations struct {
		Termination `xml:"termination"`
	} `xml:"terminations"`
	Penalties struct {
		Penalty   `xml:"penalty"`
		PrintForm zfcs_contract_printFormType `xml:"printForm"`
	} `xml:"penalties"`
	PrintForm              zfcs_contract_printFormType      `xml:"printForm"`              /* Печатная форма информации об исполнении (расторжении) контракта */
	PaymentDocuments       zfcs_contract_attachmentListType `xml:"paymentDocuments"`       /* Документы, подтверждающие исполнение, оплату */
	ReceiptDocuments       zfcs_contract_attachmentListType `xml:"receiptDocuments"`       /* Документы, подтверждающие приемку товара */
	ProductOriginDocuments zfcs_contract_attachmentListType `xml:"productOriginDocuments"` /* Информация о стране происхождения товара */
	Reason                 string                           `xml:"reason"`                 /* Основание изменения (исправления) опубликованной редакции. Указывается в случае передачи измененных сведений */
	CurrentContractStage   zfcs_contract_contractStageType  `xml:"currentContractStage"`   /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
}

type zfcs_contractProcedure2015Type struct {
	Id            int64                        `xml:"id"`            /* Идентификатор документа ЕИС */
	ExternalId    zfcs_externalIdType          `xml:"externalId"`    /* Внешний идентификатор документа */
	RegNum        zfcs_contract_regNum2015Type `xml:"regNum"`        /* Номер реестровой записи государственного или муниципального контракта */
	PublishDate   string                       `xml:"publishDate"`   /* Дата публикации */
	VersionNumber int                          `xml:"versionNumber"` /* Номер редакции сведений */
	Executions    struct {
		Stage               `xml:"stage"`
		OrdinalNumber       int  `xml:"ordinalNumber"`
		FinalStageExecution bool `xml:"finalStageExecution"`
		Execution           `xml:"execution"`
		ProductsCountries   `xml:"productsCountries"`
	} `xml:"executions"`
	Termination struct {
		Paid            zfcs_moneyPositiveType             `xml:"paid"`
		TerminationDate string                             `xml:"terminationDate"`
		ReasonInfo      zfcs_longTextMinType               `xml:"reasonInfo"`
		Reason          zfcs_contractTerminationReasonType `xml:"reason"`
		DocTermination  `xml:"docTermination"`
		DecisionDate    string `xml:"decisionDate"`
		Reparations     `xml:"reparations"`
	} `xml:"termination"`
	BankGuaranteeTermination zfcs_contractProcedure2015BankGuaranteeTerminationType `xml:"bankGuaranteeTermination"` /* Информация о прекращении обязательств поставщика, обеспеченных банковской гарантией */
	Penalties                struct {
		PenaltyAccrual `xml:"penaltyAccrual"`
		PenaltyReturn  `xml:"penaltyReturn"`
	} `xml:"penalties"`
	DelayWriteOffPenalties struct {
		TotalAmount             zfcs_moneyPositiveType `xml:"totalAmount"`
		DelayPenaltiesInProcent zfcs_valueType         `xml:"delayPenaltiesInProcent"`
		DelayPenalties          `xml:"delayPenalties"`
		WriteOffPenalties       `xml:"writeOffPenalties"`
	} `xml:"delayWriteOffPenalties"`
	BankGuaranteePayment struct {
		RegNumber                    zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`
		DocNumber                    zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`
		ImproperExecInfo             zfcs_longTextMinType            `xml:"improperExecInfo"`
		Requirements                 `xml:"requirements"`
		Paid                         `xml:"paid"`
		BankCancelDetails            string               `xml:"bankCancelDetails"`
		ImproperGuaranteePaymentInfo zfcs_longTextMinType `xml:"improperGuaranteePaymentInfo"`
		Restructure                  `xml:"restructure"`
	} `xml:"bankGuaranteePayment"`
	HoldCashEnforcement struct {
		ImproperSupplierInfo zfcs_longTextMinType          `xml:"improperSupplierInfo"`
		Currency             zfcs_currencyRef              `xml:"currency"`
		HoldAmount           zfcs_moneyPositiveType        `xml:"holdAmount"`
		HoldDate             string                        `xml:"holdDate"`
		CurrencyRate         zfcs_currencyRateContract2015 `xml:"currencyRate"`
		HoldAmountRUR        zfcs_moneyPositiveType        `xml:"holdAmountRUR"`
	} `xml:"holdCashEnforcement"`
	PrintForm              zfcs_contract_printFormType      `xml:"printForm"`              /* Печатная форма информации об исполнении (расторжении) контракта */
	ExtPrintForm           zfcs_extPrintFormType            `xml:"extPrintForm"`           /* Электронный документ, полученный из внешней системы */
	PaymentDocuments       zfcs_contract_attachmentListType `xml:"paymentDocuments"`       /* Документы, подтверждающие исполнение, оплату */
	ReceiptDocuments       zfcs_contract_attachmentListType `xml:"receiptDocuments"`       /* Документы, подтверждающие приемку товара */
	ProductOriginDocuments zfcs_contract_attachmentListType `xml:"productOriginDocuments"` /* Информация о стране происхождения товара */
	ModificationReason     zfcs_longTextMinType             `xml:"modificationReason"`     /* Основание изменения (исправления) опубликованной редакции. Указывается в случае передачи измененных сведений */
	CurrentContractStage   zfcs_contract_contractStageType  `xml:"currentContractStage"`   /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
	Okpd2okved2 bool `xml:"okpd2okved2"` /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
}

type zfcs_contractProcedure2015BankGuaranteeTerminationType struct {
	RegNumber         zfcs_bankGuaranteeRegNumberType `xml:"regNumber"`         /* Номер реестровой записи банковской гарантии */
	DocNumber         zfcs_bankGuaranteeDocNumberType `xml:"docNumber"`         /* Номер документа реестровой записи банковской гарантии (для печатной формы) */
	TerminationDate   string                          `xml:"terminationDate"`   /* Дата прекращения обязательств поставщика, обеспеченных банковской гарантией */
	TerminationReason zfcs_longTextMinType            `xml:"terminationReason"` /* Основание прекращения обязательств поставщика, обеспеченных банковской гарантией */
}

type zfcs_contractSignType struct {
	Id         int64                    `xml:"id"`       /* Идентификатор документа ЕИС */
	Number     zfcs_contract_NumberType `xml:"number"`   /* Номер контракта */
	SignDate   string                   `xml:"signDate"` /* Дата заключения контракта */
	Foundation struct {
		Order `xml:"order"`
	} `xml:"foundation"`
	Customer zfcs_organizationRef `xml:"customer"` /* Заказчик */
	Price    zfcs_moneyType       `xml:"price"`    /* Цена контракта
	(в валюте контракта) */
	PriceRUR              zfcs_moneyType   `xml:"priceRUR"`              /* Цена контракта (в рублях) */
	Currency              zfcs_currencyRef `xml:"currency"`              /* Валюта контракта */
	ConcludeContractRight bool             `xml:"concludeContractRight"` /* Признак аукциона на право заключить контракт */
	ProtocolDate          string           `xml:"protocolDate"`          /* Дата подведения итогов электронного аукциона */
	Suppliers             struct {
		Supplier zfcs_participantType `xml:"supplier"`
	} `xml:"suppliers"`
	Document struct {
		PublishedContentId       zfcs_guidType                  `xml:"publishedContentId"`
		DocumentType             zfcs_documentType70ArticleType `xml:"documentType"`
		FileName                 string                         `xml:"fileName"`
		FileSize                 string                         `xml:"fileSize"`
		DocDescription           string                         `xml:"docDescription"`
		CustomerSignature        `xml:"customerSignature"`
		SupplierSignature        `xml:"supplierSignature"`
		ControlPersonalSignature `xml:"controlPersonalSignature"`
	} `xml:"document"`
}

type zfcs_contractType struct {
	Id            int64                    `xml:"id"`            /* Идентификатор документа ЕИС */
	ExternalId    zfcs_externalIdType      `xml:"externalId"`    /* Внешний идентификатор документа */
	RegNum        zfcs_contract_regNumType `xml:"regNum"`        /* Номер реестровой записи */
	Number        zfcs_contract_NumberType `xml:"number"`        /* Номер контракта */
	PublishDate   string                   `xml:"publishDate"`   /* Дата публикации */
	SignDate      string                   `xml:"signDate"`      /* Дата заключения контракта */
	VersionNumber int                      `xml:"versionNumber"` /* Номер изменения */
	Foundation    `xml:"foundation"`       /* Основание заключения контракта */
	Customer      `xml:"customer"`         /* Заказчик */
	ProtocolDate  string                   `xml:"protocolDate"` /* Дата подведения результатов определения поставщика  (подрядчика, исполнителя) */
	DocumentBase  string                   `xml:"documentBase"` /* Реквизиты документа, подтверждающего основание заключения контракта */
	Price         zfcs_moneyType           `xml:"price"`        /* Цена контракта
	(в рублях) */
	Currency             zfcs_currencyRef `xml:"currency"` /* Валюта контракта */
	SingleCustomerReason struct {
		Id   int64                                `xml:"id"`
		Name zfcs_nsiSingleCustomerReasonNameType `xml:"name"`
	} `xml:"singleCustomerReason"`
	PriceChangeReason struct {
		Id      int64  `xml:"id"`
		Name    string `xml:"name"`
		Comment string `xml:"comment"`
	} `xml:"priceChangeReason"`
	ExecutionDate zfcs_stageType `xml:"executionDate"` /* Срок исполнения */
	Finances      struct {
		FinanceSource  string `xml:"financeSource"`
		Budget         `xml:"budget"`
		BudgetLevel    zfcs_budgetLevelType `xml:"budgetLevel"`
		Budgetary      `xml:"budgetary"`
		Extrabudget    `xml:"extrabudget"`
		Extrabudgetary `xml:"extrabudgetary"`
	} `xml:"finances"`
	Products struct {
		Product `xml:"product"`
	} `xml:"products"`
	Suppliers struct {
		Supplier zfcs_participantType `xml:"supplier"`
	} `xml:"suppliers"`
	Href             zfcs_hrefType                    `xml:"href"`             /* Гиперссылка на опубликованные сведения о контракте */
	PrintForm        zfcs_contract_printFormType      `xml:"printForm"`        /* Печатная форма контракта */
	ScanDocuments    zfcs_contract_attachmentListType `xml:"scanDocuments"`    /* Отсканированная копия контракта */
	MedicalDocuments zfcs_contract_attachmentListType `xml:"medicalDocuments"` /* Документы решения врачебной комиссии */
	Attachments      zfcs_contract_attachmentListType `xml:"attachments"`      /* Информация о прикрепленных документах */
	Modification     struct {
		Type        zfcs_contractModificationType `xml:"type"`
		Description string                        `xml:"description"`
		Base        string                        `xml:"base"`
	} `xml:"modification"`
	CurrentContractStage zfcs_contract_contractStageType `xml:"currentContractStage"` /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
}

type zfcs_contract2015Type struct {
	Id            int64               `xml:"id"`            /* Идентификатор документа */
	ExternalId    zfcs_externalIdType `xml:"externalId"`    /* Внешний идентификатор документа */
	PublishDate   string              `xml:"publishDate"`   /* Дата публикации */
	VersionNumber int                 `xml:"versionNumber"` /* Номер изменения */
	Foundation    `xml:"foundation"`  /* Основание заключения контракта */
	Customer      `xml:"customer"`    /* Заказчик */
	Placer        struct {
		ResponsibleOrg  zfcs_organizationRef            `xml:"responsibleOrg"`
		ResponsibleRole zfcs_placerRoleContract2015Type `xml:"responsibleRole"`
	} `xml:"placer"`
	Finances struct {
		BudgetFunds      `xml:"budgetFunds"`
		ExtrabudgetFunds `xml:"extrabudgetFunds"`
	} `xml:"finances"`
	ProtocolDate string                       `xml:"protocolDate"` /* Дата подведения результатов определения поставщика  (подрядчика, исполнителя) */
	DocumentBase zfcs_longTextMinType         `xml:"documentBase"` /* Реквизиты документа, подтверждающего основание заключения контракта */
	DocumentCode string                       `xml:"documentCode"` /* Код документа, подтверждающего основание заключения контракта (для печатной формы) */
	SignDate     string                       `xml:"signDate"`     /* Дата заключения контракта */
	RegNum       zfcs_contract_regNum2015Type `xml:"regNum"`       /* Номер реестровой записи */
	Number       zfcs_contract_Number2015Type `xml:"number"`       /* Номер контракта */
	PriceInfo    struct {
		Price        zfcs_moneyPositiveType        `xml:"price"`
		PriceType    string                        `xml:"priceType"`
		PriceFormula zfcs_longTextMinType          `xml:"priceFormula"`
		Currency     zfcs_currencyRef              `xml:"currency"`
		CurrencyRate zfcs_currencyRateContract2015 `xml:"currencyRate"`
		PriceRUR     zfcs_moneyPositiveType        `xml:"priceRUR"`
	} `xml:"priceInfo"`
	SubContractorsSum struct {
		SumInPercents float64                `xml:"sumInPercents"`
		PriceValueRUR zfcs_moneyPositiveType `xml:"priceValueRUR"`
	} `xml:"subContractorsSum"`
	ExecutionPeriod struct {
		StartDate string `xml:"startDate"`
		Stages    `xml:"stages"`
		EndDate   string `xml:"endDate"`
	} `xml:"executionPeriod"`
	Enforcement               `xml:"enforcement"`                      /* Обеспечение исполнения контракта */
	GuaranteeReturns          zfcs_contract2015BankGuaranteeReturnType `xml:"guaranteeReturns"`          /* Информация о возвращении банковской гарантии или уведомление об освобождении от обязательств по банковской гарантии */
	EnergyServiceContractInfo zfcs_longTextMinType                     `xml:"energyServiceContractInfo"` /* Информация об экономии при заключении энергосервисного контракта */
	Products                  struct {
		Product        `xml:"product"`
		ProductsChange `xml:"productsChange"`
	} `xml:"products"`
	Suppliers struct {
		Supplier zfcs_contract2015SupplierType `xml:"supplier"`
	} `xml:"suppliers"`
	Href             zfcs_hrefType                    `xml:"href"`             /* Гиперссылка на опубликованные сведения о контракте */
	PrintForm        zfcs_contract_printFormType      `xml:"printForm"`        /* Печатная форма контракта */
	ExtPrintForm     zfcs_extPrintFormType            `xml:"extPrintForm"`     /* Электронный документ, полученный из внешней системы */
	ScanDocuments    zfcs_contract_attachmentListType `xml:"scanDocuments"`    /* Отсканированная копия контракта */
	MedicalDocuments zfcs_contract_attachmentListType `xml:"medicalDocuments"` /* Документы решения врачебной комиссии */
	Attachments      zfcs_contract_attachmentListType `xml:"attachments"`      /* Информация о прикрепленных документах */
	Modification     struct {
		Attachments zfcs_contract_attachmentListType `xml:"attachments"`
	} `xml:"modification"`
	CurrentContractStage zfcs_contract_contractStageType `xml:"currentContractStage"` /* Текущее состояние контракта:
	E 	Исполнение;
	ET	Исполнение прекращено;
	EC	Исполнение завершено;
	IN	Aннулировано. */
	Okpd2okved2 bool `xml:"okpd2okved2"` /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
}

type zfcs_contract2015BankGuaranteeReturnType struct {
	GuaranteeReturn `xml:"guaranteeReturn"` /* Информация */
}

type zfcs_extraBudgetFundsContract2015 struct {
	Code string `xml:"code"` /* Код вида внебюджетных средств */
	Name string `xml:"name"` /* Наименование вида внебюджетных средств */
}

type zfcs_currencyRateContract2015 struct {
	Rate    float64 `xml:"rate"`    /* Курс валюты по отношению к рублю */
	Raiting int     `xml:"raiting"` /* Номинал валюты */
}

type zfcs_budgetFundsContract2015 struct {
	Code string               `xml:"code"` /* Код бюджета */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование бюджета */
}

type zfcs_stageType struct {
	Month zfcs_monthType `xml:"month"` /* Месяц окончания этапа */
	Year  zfcs_yearType  `xml:"year"`  /* Год окончания этапа */
}

type zfcs_subStageType struct {
	BasedOn       zfcs_stageType
	SubstageMonth zfcs_monthType `xml:"substageMonth"` /* Месяц финансирования в этапе */
	SubstageYear  zfcs_yearType  `xml:"substageYear"`  /* Год финансирования в этапе */
}

type zfcs_standardContractType struct {
	Id                     int64                           `xml:"id"`                     /* Идентификатор документа */
	DocPublishDate         string                          `xml:"docPublishDate"`         /* Дата публикации документа */
	StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракт, типовых условий контракта */
	Type                   zfcs_standardContractTypeEnum   `xml:"type"`                   /* Вид документа:
	С - типовой контракт,
	Т - типовые условия контракта */
	ApproveInfo struct {
		Organization zfcs_purchaseOrganizationType `xml:"organization"`
		Date         string                        `xml:"date"`
		Document     `xml:"document"`
	} `xml:"approveInfo"`
	PlacerOrganization zfcs_purchaseOrganizationType `xml:"placerOrganization"` /* Организация, разместившая сведения */
	Indications        struct {
		PurchaseObjects `xml:"purchaseObjects"`
		ContractPrice   `xml:"contractPrice"`
		OtherIndicators zfcs_longTextMinType `xml:"otherIndicators"`
	} `xml:"indications"`
	UseCases struct {
		Type          string               `xml:"type"`
		UseTerms      zfcs_longTextMinType `xml:"useTerms"`
		AddInfo       zfcs_longTextMinType `xml:"addInfo"`
		RequiredTerms zfcs_longTextMinType `xml:"requiredTerms"`
	} `xml:"useCases"`
	PrintForm    zfcs_printFormType      `xml:"printForm"`   /* Печатная форма контракта */
	Attachments  zfcs_attachmentListType `xml:"attachments"` /* Документация контракта */
	Modification struct {
		Version int               `xml:"version"`
		Info    zfcs_longTextType `xml:"info"`
	} `xml:"modification"`
	Okpd2okved2 bool `xml:"okpd2okved2"` /* Классификация по ОКПД2/ОКВЭД2 */
}

type zfcs_standardContractInvalidType struct {
	Id                     int64                           `xml:"id"`                     /* Идентификатор документа ЕИС */
	DocPublishDate         string                          `xml:"docPublishDate"`         /* Дата публикации документа */
	StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракт, типовых условий контракта */
	PlacerOrganization     zfcs_purchaseOrganizationType   `xml:"placerOrganization"`     /* Организация, разместившая сведения */
	Type                   zfcs_standardContractTypeEnum   `xml:"type"`                   /* Вид документа:
	С - типовой контракт,
	Т - типовые условия контракта */
	ApproveInfo struct {
		Organization zfcs_purchaseOrganizationType `xml:"organization"`
		Date         string                        `xml:"date"`
		Document     `xml:"document"`
	} `xml:"approveInfo"`
	Indications struct {
		PurchaseObjects `xml:"purchaseObjects"`
		ContractPrice   `xml:"contractPrice"`
		OtherIndicators zfcs_longTextMinType `xml:"otherIndicators"`
	} `xml:"indications"`
	UseCases struct {
		Type          string               `xml:"type"`
		UseTerms      zfcs_longTextMinType `xml:"useTerms"`
		AddInfo       zfcs_longTextMinType `xml:"addInfo"`
		RequiredTerms zfcs_longTextMinType `xml:"requiredTerms"`
	} `xml:"useCases"`
	Okpd2okved2    bool                    `xml:"okpd2okved2"` /* Классификация по ОКПД2/ОКВЭД2 */
	PrintForm      zfcs_printFormType      `xml:"printForm"`   /* Печатная форма контракта */
	Attachments    zfcs_attachmentListType `xml:"attachments"` /* Документация контракта */
	InvalidityInfo struct {
		Date   string            `xml:"date"`
		Reason zfcs_longTextType `xml:"reason"`
	} `xml:"invalidityInfo"`
}

type zfcs_customerReportBaseType struct {
	Id             int64                             `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType               `xml:"externalId"`     /* Внешний идентификатор документа */
	DocDate        string                            `xml:"docDate"`        /* Дата документа */
	DocPublishDate string                            `xml:"docPublishDate"` /* Дата публикации в ЕИС */
	VersionNumber  zfcs_revisionNumType              `xml:"versionNumber"`  /* Номер редакции */
	PublishOrg     zfcs_organizationInfoWithOgrnType `xml:"publishOrg"`     /* Информация об организации, разместившей отчет */
	Href           zfcs_hrefType                     `xml:"href"`           /* Гиперссылка на опубликованные сведения об отчете. Элемент не используется в импорте */
	PrintForm      zfcs_printFormType                `xml:"printForm"`      /* Печатная форма отчета. */
	ExtPrintForm   zfcs_extPrintFormType             `xml:"extPrintForm"`   /* Электронный документ, полученный из внешней системы */
	Attachments    zfcs_attachmentListType           `xml:"attachments"`    /* Копии документов */
}

type zfcs_customerReportType struct {
	BasedOn            zfcs_customerReportBaseType
	ReportId           zfcs_reportIdType    `xml:"reportId"`           /* Номер отчета в ЕИС */
	ModificationReason zfcs_longTextMinType `xml:"modificationReason"` /* Причина редактирования информации */
	Signer             struct {
		FirstName  string `xml:"firstName"`  /* Имя */
		MiddleName string `xml:"middleName"` /* Отчество */
		LastName   string `xml:"lastName"`   /* Фамилия */
		Position   string `xml:"position"`   /* Должность */
	} `xml:"signer"`
}

type zfcs_customerReportInvalidType struct {
	BasedOn           zfcs_customerReportBaseType
	ReportId          zfcs_reportIdType      `xml:"reportId"`          /* Номер отчета в ЕИС */
	InvalidReportInfo zfcs_invalidReportType `xml:"invalidReportInfo"` /* Информация о недействительности сведений отчета */
}

type zfcs_customerReportContractExecutionType struct {
	BasedOn      zfcs_customerReportType
	Customer     `xml:"customer"` /* Информация о заказчике */
	ContractInfo struct {
		ContractRegNum zfcs_contract_regNumType `xml:"contractRegNum"` /* Реестровый номер контракта */
		ContractNumber zfcs_contract_NumberType `xml:"contractNumber"` /* Номер контракта. Элемент не используется в импорте */
		Product        `xml:"product"`          /* Предмет контракта */
		PurchaseCode   zfcs_ikzCodeType         `xml:"purchaseCode"` /* Идентификационный код закупки. Элемент не используется в импорте */
		Okpd2okved2    bool                     `xml:"okpd2okved2"`  /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
	} `xml:"contractInfo"`
	Suppliers struct {
		Supplier `xml:"supplier"` /* Поставщик */
	} `xml:"suppliers"`
	ExecutionInfo struct {
		ProvideByContract zfcs_contractExecutionType `xml:"provideByContract"` /* Дата начала исполнения контракта (отдельного этапа исполнения контракта) */
		Executed          zfcs_contractExecutionType `xml:"executed"`          /* Дата окончания исполнения контракта (отдельного этапа исполнения контракта) */
		Price             `xml:"price"`              /* Цена контракта (отдельного этапа исполнения контракта): в валюте контракта в рублях */
		Quantity          `xml:"quantity"`           /* Количество (объем) поставляемых товаров, оказываемых услуг, выполняемых работ по контракту (по отдельному этапу исполнения контракта) с указанием единиц измерения (в случае если объект закупки может быть количественно измерен) */
	} `xml:"executionInfo"`
	ImproperExecutionInfo `xml:"improperExecutionInfo"` /* Информация о ненадлежащем исполнении контракта или неисполнении контракта и о санкциях, связанных с указанным нарушением */
	ModifyTerminateInfo   struct {
		ModifyContract    `xml:"modifyContract"`    /* Изменение контракта */
		TerminateContract `xml:"terminateContract"` /* Расторжение контракта */
	} `xml:"modifyTerminateInfo"`
}

type zfcs_customerReportContractExecutionInvalidType struct {
	BasedOn zfcs_customerReportInvalidType
	Report  zfcs_customerReportContractExecutionType `xml:"report"` /* Информация отчета. Элемент не используется в импорте */
}

type zfcs_customerReportSmallScaleBusinessType struct {
	BasedOn          zfcs_customerReportType
	Customer         zfcs_organizationInfoExtendedType `xml:"customer"`        /* Информация о заказчике */
	ReportingPeriod  zfcs_yearType                     `xml:"reportingPeriod"` /* Отчетный период */
	QuantityPurchase struct {
		Privacy          zfcs_volumeType `xml:"privacy"`          /* 1.1. Совокупный годовой объем закупок, за исключением объема закупок, сведения о которых составляют государственную тайну */
		AnnualVolumeSt30 zfcs_volumeType `xml:"annualVolumeSt30"` /* 1.2. Общий объем финансового обеспечения для оплаты контрактов в отчетном году в рамках осуществления закупок, предусмотренных частью 1.1 статьи 30 Федерального закона "О контрактной системе в сфере закупок товаров, работ, услуг для государственных и муниципальных нужд" (далее - Федеральный закон) (тыс. руб.). Элемент не используется в импорте
		, значение рассчитывается по сумме последующих полей */
		Lending zfcs_volumeType `xml:"lending"` /* 1.2.1. Объем финансового обеспечения для оплаты в отчетном году контрактов, заключаемых на оказание услуг по предоставлению кредитов
		(тыс. руб.) */
		SingleSupplier zfcs_volumeType `xml:"singleSupplier"` /* 1.2.2. Объем финансового обеспечения для оплаты в отчетном году контрактов, заключаемых с единственным поставщиком (подрядчиком, исполнителем)
		в соответствии с частью 1 статьи 93  Федерального закона (тыс. руб.) */
		NuclearEnergy zfcs_volumeType `xml:"nuclearEnergy"` /* 1.2.3. Объем финансового обеспечения для оплаты в отчетном году контрактов, заключаемых на выполнение работ в области использования атомной энергии (тыс. руб.) */
		ZK            zfcs_volumeType `xml:"ZK"`            /* 1.2.4 Объем финансового обеспечения для оплаты в отчетном году контрактов, заключаемых по результатам закрытых способов определения поставщиков (подрядчиков, исполнителей) (тыс. руб.) */
		AnnualVolume  zfcs_volumeType `xml:"annualVolume"`  /* 1.3. Совокупный годовой объем закупок, рассчитанный за вычетом закупок, предусмотренных частью 1.1 статьи 30 Федерального закона (тыс. руб.). Элемент не используется в импорте
		, значение рассчитывается по разнице полей 1.2 и 1.1 */
		Percent15       zfcs_volumeType `xml:"percent15"`       /* 1.4. Объем закупок, который заказчик обязан осуществить у СМП, СОНО в отчетном году (не менее чем 15 процентов совокупного годового объема закупок, рассчитанного с учетом части 1.1   статьи 30 Федерального закона) (тыс. руб.) */
		OnlySMP         zfcs_volumeType `xml:"onlySMP"`         /* 2.1. Объем закупок в отчетном году, осуществленных по результатам состоявшихся процедур определения поставщика (подрядчика, исполнителя), в извещении об осуществлении которых было установлено ограничение в отношении участников закупок, которыми могли быть только СМП, СОНО (тыс. руб.) */
		NotSMP          zfcs_volumeType `xml:"notSMP"`          /* 2.2. Объем привлечения в отчетном году субподрядчиков, соисполнителей из числа СМП, СОНО к исполнению контрактов, заключенных по результатам определений поставщиков (подрядчиков, исполнителей), в извещениях об осуществлении которых было установлено требование к поставщику (подрядчику, исполнителю), не являющемуся СМП, СОНО о привлечении к исполнению контракта субподрядчиков, соисполнителей из числа СМП, СОНО (тыс. руб.) */
		AnnualVolumeSMP zfcs_volumeType `xml:"annualVolumeSMP"` /* 2.3. Объем закупок, который заказчик осуществил у СМП, СОНО в отчетном году (тыс. руб.). Элемент не используется в импорте
		, значение рассчитывается по сумме полей 2.1 и 2.2 */
		RateSMP zfcs_volumeType `xml:"rateSMP"` /* 2.4. Доля закупок, которые заказчик осуществил у СМП, СОНО в отчетном году (%)/ Элемент не используется в импорте
		, значение рассчитывается по отношению полей 2.3 к 1.3 */
		AbandonedSum zfcs_volumeType `xml:"abandonedSum"` /* 3.1. Сумма начальных (максимальных) цен контрактов несостоявшихся определений поставщиков (подрядчиков, исполнителей) с участием СМП, СОНО (тыс. руб.) */
	} `xml:"quantityPurchase"`
	ContractsInfo struct {
		ContractInfo `xml:"contractInfo"` /* Информация о заключенном контракте */
	} `xml:"contractsInfo"`
}

type zfcs_customerReportSmallScaleBusinessInvalidType struct {
	BasedOn zfcs_customerReportInvalidType
	Report  zfcs_customerReportSmallScaleBusinessType `xml:"report"` /* Информация отчета. Элемент не используется в импорте */
}

type zfcs_customerReportSingleContractorType struct {
	BasedOn        zfcs_customerReportType
	Customer       zfcs_organizationInfoWithOgrnType `xml:"customer"`       /* Информация о заказчике */
	PurchaseNumber zfcs_purchaseNumberType           `xml:"purchaseNumber"` /* Номер извещения об осуществлении закупки */
	ContractRegNum zfcs_contract_regNumType          `xml:"contractRegNum"` /* Реестровый номер контракта */
	LotNumber      int                               `xml:"lotNumber"`      /* Номер лота в извещении */
	Reason         struct {
		Code zfcs_nsiSingleCustomerReasonCodeType `xml:"code"` /* Код */
		Name zfcs_nsiSingleCustomerReasonNameType `xml:"name"` /* Наименование основания */
	} `xml:"reason"`
}

type zfcs_customerReportSingleContractorInvalidType struct {
	BasedOn zfcs_customerReportInvalidType
	Report  zfcs_customerReportSingleContractorType `xml:"report"` /* Информация отчета. Элемент не используется в импорте */
}

type zfcs_customerReportBigProjectMonitoringType struct {
	BasedOn        zfcs_customerReportType
	NeedInContract bool                              `xml:"needInContract"` /* Для реализации инвестиционного проекта должен заключаться государственный контракт */
	Customer       zfcs_organizationInfoWithOgrnType `xml:"customer"`       /* Информация о заказчике */
	ContractsInfo  struct {
		ContractInfo `xml:"contractInfo"` /* Информация о заключенном контракте */
	} `xml:"contractsInfo"`
	Constructor struct {
		FullName  zfcs_longTextMinType `xml:"fullName"`  /* Полное наименование */
		ShortName zfcs_longTextMinType `xml:"shortName"` /* Сокращенное наименование */
		OKOPF     zfcs_okopfRef        `xml:"OKOPF"`     /* ОКОПФ организации */
		Head      zfcs_longTextMinType `xml:"head"`      /* Ф.И.О. руководителя */
		Position  zfcs_longTextMinType `xml:"position"`  /* Должность руководителя  */
	} `xml:"constructor"`
	ProjectInfo struct {
		Name      zfcs_longTextMinType `xml:"name"`    /* Наименование инвестиционного проекта */
		Purpose   zfcs_longTextMinType `xml:"purpose"` /* Цель инвестиционного проекта */
		Period    `xml:"period"`       /* Срок реализации инвестиционного проекта */
		Direction string               `xml:"direction"` /* Направление реализации инвестиционного проекта.
		Допустимые значения:
		IT1 - строительство,
		IT2 - реконструкция объекта капитального строительства,
		IT3 - иные инвестиции в основной капитал */
		OtherDirection zfcs_longTextMinType `xml:"otherDirection"` /* Иной вид инвестиции в основной капитал. Указывается если выбран 3 тип направления реализации инвестиционного проекта */
		Mechanism      string               `xml:"mechanism"`      /* Механизм, в рамках которого представляются средства федерального бюджета.
		Допустимые значения:
		FP1 - федеральная инвестиционная программа,
		FP2 - Инвестиционный фонд Российской Федерации,
		FP3 - Фонд национального благосостояния,
		FP4 - иной механизм с его указанием */
		OtherMechanism zfcs_longTextMinType `xml:"otherMechanism"` /* Иной механизм, в рамках которого представляются средства федерального бюджета. Указывается если выбран 4 тип направления реализации инвестиционного проекта */
		Grbs           zfcs_longTextMinType `xml:"grbs"`           /* Главный распорядитель средств федерального бюджета */
		Indicators     `xml:"indicators"`   /* Количественные показатели (показатель) результатов реализации инвестиционного проекта */
	} `xml:"projectInfo"`
	Contractors struct {
		Contractor zfcs_bigProjectMemberType `xml:"contractor"` /* Субподрядчик */
	} `xml:"contractors"`
	Participants struct {
		Participant zfcs_bigProjectMemberType `xml:"participant"` /* Участник */
	} `xml:"participants"`
	Documents struct {
		ProjectDocumentation `xml:"projectDocumentation"` /* Проектная документация по инвестиционному проекту */
		Statement            `xml:"statement"`            /* Положительное заключение государственной экспертизы проектной документации и результаты инженерных изысканий */
		Examination          `xml:"examination"`          /* Положительное заключение о достоверности сметной стоимости инвестиционного проекта (объекта) */
		Audit                `xml:"audit"`                /* Положительное заключение о проведении публичного технологического и ценового аудита инвестиционного проекта, в случае если его проведение предусмотрено законодательством Российской Федерации */
	} `xml:"documents"`
	Cost struct {
		Act          `xml:"act"`            /* В соответствии с актом о предоставлении средств федерального бюджета на реализацию инвестиционного проекта (в случае если принятие такого акту предусмотрено Российской Федерацией), актом об утверждении паспорта инвестиционного проекта, соглашением о предоставлении средств федерального бюджета, рассчитанная в ценах соответствующих лет */
		Authenticity `xml:"authenticity"`   /* В соответствии с заключением о достоверности сметной стоимости инвестиционного проекта (объекта) в ценах года ее определения */
		Contract     `xml:"contract"`       /* По результатам заключения государственного контракта (договора на подрядные работы) по инвестиционному проекту с указанием сроков его исполнения  */
		Economy      zfcs_moneyPositiveType `xml:"economy"` /* Экономия по результатам проведения конкурсных процедур по заключению государственного контракта (договора) на реализацию инвестиционного проекта, в млн. рублей */
	} `xml:"cost"`
	Financings struct {
		Years  zfcs_bigProjectFinancingYearsType `xml:"years"` /* Годы реализации инвестиционного проекта  */
		Stages `xml:"stages"`                    /* Этапы реализации инвестиционного проекта */
		Site   zfcs_longTextMinType              `xml:"site"` /* Адрес сайта в информационно-телекоммуникационной сети "Интернет", на котором размещена информация о реализации инвестиционного проекта  */
	} `xml:"financings"`
	Tenders struct {
		Tender `xml:"tender"` /* Конкурсная процедура */
		Total  `xml:"total"`  /* Итого */
	} `xml:"tenders"`
	Realization struct {
		Year `xml:"year"` /* Финансовый год */
	} `xml:"realization"`
	Changes struct {
		Year `xml:"year"` /* Финансовый год */
	} `xml:"changes"`
}

type zfcs_customerReportBigProjectMonitoringInvalidType struct {
	BasedOn zfcs_customerReportInvalidType
	Report  zfcs_customerReportBigProjectMonitoringType `xml:"report"` /* Информация отчета. Элемент не используется в импорте */
}

type zfcs_contractExecutionType struct {
	ContractDate  string               `xml:"contractDate"`  /* Дата предусмотренная контрактом */
	ExecutionDate string               `xml:"executionDate"` /* Дата исполнения  */
	Note          zfcs_longTextMinType `xml:"note"`          /* Примечание */
}

type zfcs_improperContractExecutionType struct {
	NameObligations  zfcs_longTextMinType `xml:"nameObligations"`  /* Наименование обязательства */
	EssenceViolation zfcs_longTextMinType `xml:"essenceViolation"` /* Суть нарушения */
	PenaltyInfo      zfcs_longTextMinType `xml:"penaltyInfo"`      /* Информация о начисленной неустойке */
	PenaltyDoc       zfcs_longTextMinType `xml:"penaltyDoc"`       /* Докмент, подтверждающий начисление или уплату неустойки */
	Note             zfcs_longTextMinType `xml:"note"`             /* Примечание */
}

type zfcs_invalidReportType struct {
	PublishDate string                  `xml:"publishDate"` /* Дата публикации. Элеменрт не используется в импорте */
	Reason      zfcs_longTextType       `xml:"reason"`      /* Причина недействительности сведений */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Документы */
}

type zfcs_modifyTerminateContractType struct {
	EventDate string               `xml:"eventDate"` /* Дата события */
	DocInfo   zfcs_longTextMinType `xml:"docInfo"`   /* Наиенование, номер и дата документа-основания */
}

type zfcs_organizationInfoType struct {
	BasedOn   zfcs_organizationRef
	ShortName string                   `xml:"shortName"` /* Краткое наименование. Элемент не используется в импорте */
	INN       zfcs_innOrganizationType `xml:"INN"`       /* ИНН организации. Элемент не используется в импорте */
	KPP       zfcs_kppType             `xml:"KPP"`       /* КПП организации. Элемент не используется в импорте */
}

type zfcs_organizationInfoWithOgrnType struct {
	BasedOn         zfcs_organizationRef
	ShortName       string                   `xml:"shortName"`       /* Краткое наименование. Элемент не используется в импорте */
	INN             zfcs_innType             `xml:"INN"`             /* ИНН организации. Элемент не используется в импорте */
	KPP             zfcs_kppType             `xml:"KPP"`             /* КПП организации. Элемент не используется в импорте */
	OGRN            zfcs_ogrnType            `xml:"OGRN"`            /* ОГРН организации. Элемент не используется в импорте */
	ResponsibleRole zfcs_responsibleRoleType `xml:"responsibleRole"` /* Роль организации, осуществляющей закупку. Элемент не используется в импорте:

	CU - Заказчик;
	OCU - Заказчик в качестве организатора совместного аукциона;
	RA - Уполномоченный орган;
	ORA- Уполномоченный орган в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
	AI - Уполномоченное учреждение;
	OAI- Уполномоченное учреждение в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
	OA - Организация, осуществляющая полномочия заказчика на осуществление закупок на основании договора (соглашения);
	OOA- Организация, осуществляющая полномочия заказчика на осуществление закупок на основании договора (соглашения) в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
	CS - Заказчик, осуществляющий закупки в соответствии с частью 5 статьи 15 Федерального закона № 44-ФЗ;
	OCS -  Заказчик, осуществляющий закупки в соответствии с частью 5 статьи 15 Федерального закона № 44-ФЗ, в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
	CC - Заказчик, осуществляющий закупки в соответствии с Федеральным законом № 44-ФЗ, в связи с неразмещением положения о закупке в соответствии с положениями Федерального закона № 223-ФЗ;
	OCC - Заказчик, осуществляющий закупки в соответствии с Федеральным законом № 44-ФЗ, в связи с неразмещением положения о закупке в соответствии с положениями Федерального закона № 223-ФЗ,
	в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
	AU - Заказчик, осуществляющий закупку на проведение обязательного аудита (код AU);
	OAU - Заказчик, осуществляющий закупку на проведение обязательного аудита (код AU), в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ. */
}

type zfcs_organizationInfoExtendedType struct {
	BasedOn zfcs_organizationInfoType
	Adress  zfcs_longTextType `xml:"adress"` /* Местонахождение (адрес) */
	Phone   string            `xml:"phone"`  /* Телефон */
	Email   string            `xml:"email"`  /* Адрес электронной почты */
	OKTMO   zfcs_OKTMORef     `xml:"OKTMO"`  /* ОКТМО */
	OKPO    zfcs_OKPORef      `xml:"OKPO"`   /* Код по ОКПО. Элемент не используется в импорте */
	OKOPF   zfcs_okopfRef     `xml:"OKOPF"`  /* ОКОПФ организации. Элемент не используется в импорте */
}

type zfcs_bigProjectValueType struct {
	Cost       zfcs_moneyPositiveType `xml:"cost"`       /* Объем выполненных работ в млн рублей */
	Percentage float64                `xml:"percentage"` /* % от общего объема */
}

type zfcs_bigProjectFinancingYearsType struct {
	Total zfcs_bigProjectFinancingsType `xml:"total"` /* Итого */
	Year  struct {
		Financing zfcs_bigProjectFinancingsType `xml:"financing"`
		Year      zfcs_yearType                 `xml:"year"`
	} `xml:"year"`
}

type zfcs_bigProjectMemberType struct {
	Name    zfcs_longTextMinType `xml:"name"`    /* Полное наименование / Ф.И.О. */
	INN     zfcs_innAnalogType   `xml:"INN"`     /* ИНН */
	KPP     zfcs_kppType         `xml:"KPP"`     /* КПП */
	Address zfcs_longTextType    `xml:"address"` /* Местонахождение (адрес) */
}

type zfcs_bigProjectFinancingsType struct {
	Sources struct {
		Federal   zfcs_moneyPositiveType `xml:"federal"`
		Region    zfcs_moneyPositiveType `xml:"region"`
		Self      zfcs_moneyPositiveType `xml:"self"`
		Nonbudget zfcs_moneyPositiveType `xml:"nonbudget"`
	} `xml:"sources"`
	Cost zfcs_moneyPositiveType `xml:"cost"` /* Стоимость инвестиционного проекта, в млн рублей */
}

type zfcs_bigProjectCostType struct {
	Cost zfcs_longTextMinType `xml:"cost"` /* Стоимость инвестиционного проекта, рассчитанная в ценах соответствующих лет */
	Year zfcs_yearType        `xml:"year"` /* Год определения цены */
}

type zfcs_requestForQuotationType struct {
	Id                int64                 `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId        zfcs_externalIdType   `xml:"externalId"`     /* Внешний идентификатор документа */
	DocPublishDate    string                `xml:"docPublishDate"` /* Дата публикации в ЕИС */
	RegistryNum       zfcs_registryNumType  `xml:"registryNum"`    /* Реестровый номер */
	VersionNumber     zfcs_revisionNumType  `xml:"versionNumber"`  /* Номер редакции */
	State             zfcs_longTextType     `xml:"state"`          /* Статус. Элемент не используется в импорте */
	PublishOrg        `xml:"publishOrg"`    /* Информация об организации, разместившей запрос цен */
	Href              zfcs_hrefType         `xml:"href"`              /* Гиперссылка на опубликованные сведения об отчете. Элемент не используется в импорте */
	PrintForm         zfcs_printFormType    `xml:"printForm"`         /* Печатная форма отчета. */
	ExtPrintForm      zfcs_extPrintFormType `xml:"extPrintForm"`      /* Электронный документ, полученный из внешней системы */
	RequestObjectInfo zfcs_longTextType     `xml:"requestObjectInfo"` /* Наименование объекта закупки */
	ResponsibleInfo   struct {
		Place         zfcs_longTextType      `xml:"place"`
		ContactPerson zfcs_contactPersonType `xml:"contactPerson"`
		ContactEMail  zfcs_string            `xml:"contactEMail"`
		ContactPhone  zfcs_string            `xml:"contactPhone"`
		ContactFax    zfcs_string            `xml:"contactFax"`
		AddInfo       zfcs_longTextType      `xml:"addInfo"`
	} `xml:"responsibleInfo"`
	ProcedureInfo struct {
		Request  `xml:"request"`
		Purchase `xml:"purchase"`
	} `xml:"procedureInfo"`
	Products struct {
		ObjectInfo  zfcs_longTextType `xml:"objectInfo"`
		Product     `xml:"product"`
		Identity    zfcs_longTextType `xml:"identity"`
		Okpd2okved2 bool              `xml:"okpd2okved2"`
	} `xml:"products"`
	Conditions struct {
		Main              zfcs_longTextType `xml:"main"`
		Payment           zfcs_longTextType `xml:"payment"`
		ContractGuarantee zfcs_longTextType `xml:"contractGuarantee"`
		Warranty          zfcs_longTextType `xml:"warranty"`
		Delivery          zfcs_longTextType `xml:"delivery"`
		AddInfo           zfcs_longTextType `xml:"addInfo"`
	} `xml:"conditions"`
	Attachments        zfcs_attachmentListType `xml:"attachments"`        /* Документы */
	ModificationReason zfcs_longTextMinType    `xml:"modificationReason"` /* Причина редактирования информации */
}

type zfcs_requestForQuotationCancelType struct {
	Id                int64                             `xml:"id"`                /* Идентификатор документа ЕИС */
	ExternalId        zfcs_externalIdType               `xml:"externalId"`        /* Внешний идентификатор документа */
	DocPublishDate    string                            `xml:"docPublishDate"`    /* Дата публикации в ЕИС */
	RegistryNum       zfcs_registryNumType              `xml:"registryNum"`       /* Реестровый номер */
	VersionNumber     zfcs_revisionNumType              `xml:"versionNumber"`     /* Номер редакции */
	State             zfcs_longTextType                 `xml:"state"`             /* Статус. Элемент не используется в импорте */
	Href              zfcs_hrefType                     `xml:"href"`              /* Гиперссылка на опубликованные сведения об отчете. Элемент не используется в импорте */
	PrintForm         zfcs_printFormType                `xml:"printForm"`         /* Печатная форма отчета. */
	ExtPrintForm      zfcs_extPrintFormType             `xml:"extPrintForm"`      /* Электронный документ, полученный из внешней системы */
	PublishOrg        zfcs_organizationInfoWithOgrnType `xml:"publishOrg"`        /* Информация об организации, разместившей запрос цен. Элемент не используется в импорте */
	RequestObjectInfo zfcs_longTextType                 `xml:"requestObjectInfo"` /* Наименование объекта закупки. Элемент не используется в импорте */
	ResponsibleInfo   struct {
		Place         zfcs_longTextType      `xml:"place"`
		ContactPerson zfcs_contactPersonType `xml:"contactPerson"`
		ContactEMail  zfcs_string            `xml:"contactEMail"`
		ContactPhone  zfcs_string            `xml:"contactPhone"`
		ContactFax    zfcs_string            `xml:"contactFax"`
		AddInfo       zfcs_longTextType      `xml:"addInfo"`
	} `xml:"responsibleInfo"`
	ProcedureInfo struct {
		Request  `xml:"request"`
		Purchase `xml:"purchase"`
	} `xml:"procedureInfo"`
	Products struct {
		ObjectInfo  zfcs_longTextType `xml:"objectInfo"`
		Product     `xml:"product"`
		Identity    zfcs_longTextType `xml:"identity"`
		Okpd2okved2 bool              `xml:"okpd2okved2"`
	} `xml:"products"`
	Conditions struct {
		Main              zfcs_longTextType `xml:"main"`
		Payment           zfcs_longTextType `xml:"payment"`
		ContractGuarantee zfcs_longTextType `xml:"contractGuarantee"`
		Warranty          zfcs_longTextType `xml:"warranty"`
		Delivery          zfcs_longTextType `xml:"delivery"`
		AddInfo           zfcs_longTextType `xml:"addInfo"`
	} `xml:"conditions"`
	Attachments        zfcs_attachmentListType `xml:"attachments"`        /* Документы */
	ModificationReason zfcs_longTextMinType    `xml:"modificationReason"` /* Причина редактирования информации */
	CancelReason       zfcs_longTextType       `xml:"cancelReason"`       /* Основание отмены запроса цен */
}

type zfcs_clarificationRequestType struct {
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocNumber      string                  `xml:"docNumber"`      /* Номер запроса на разъяснение */
	DocDate        string                  `xml:"docDate"`        /* Дата поступления запроса на разъяснение */
	Topic          zfcs_longTextType       `xml:"topic"`          /* Тема запроса на разъяснение /
	краткое описание запроса на разъяснение */
	Participant struct {
		Name  zfcs_longTextType `xml:"name"`
		Email zfcs_string       `xml:"email"`
	} `xml:"participant"`
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
}

type zfcs_clarificationType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор объекта ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	RequestNumber  string                  `xml:"requestNumber"`  /* Номер запроса на разъяснение */
	DocNumber      string                  `xml:"docNumber"`      /* Номер разъяснения */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации разъяснения */
	Href           zfcs_hrefType           `xml:"href"`           /* Гиперссылка на опубликованное разъяснение */
	ExtPrintForm   zfcs_extPrintFormType   `xml:"extPrintForm"`   /* Электронный документ, полученный из внешней системы */
	Question       zfcs_longTextType       `xml:"question"`       /* Тема запроса на разъяснение */
	Topic          zfcs_longTextType       `xml:"topic"`          /* Тема разъяснения /
	краткое описание разъяснения */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
}

type zfcs_notificationCancelType struct {
	BasedOn      zfcs_purchaseDocumentCommonType
	CancelReason zfcs_purchaseCancelReasonType `xml:"cancelReason"` /* Причина отмены */
	AddInfo      zfcs_longTextType             `xml:"addInfo"`      /* Дополнительная информация */
}

type zfcs_notificationCancelFailureType struct {
	BasedOn                zfcs_purchaseDocumentCommonType
	NotificationCancelInfo zfcs_longTextType   `xml:"notificationCancelInfo"` /* Информация об отменяемом извещении об отмене определения поставщика  */
	PlacingWay             zfcs_placingWayType `xml:"placingWay"`             /* Подспособ определения поставщика */
	PurchaseObjectInfo     zfcs_longTextType   `xml:"purchaseObjectInfo"`     /* Наименование объекта закупки */
	Lot                    struct {
		LotNumber uint64            `xml:"lotNumber"` /* Номер лота в извещении */
		Info      zfcs_longTextType `xml:"info"`      /* Текстовое описание лотов */
	} `xml:"lot"`
	RecoveryToStage zfcs_tenderStageType `xml:"recoveryToStage"` /* Этап определения поставщика

	Допустимые значения:

	NP  Подготовка извещения;
	AP  Подача заявок;
	CW  Работа комиссии;
	FO  Определение поставщика завершено;
	CO  Определение поставщика отменено.		 */
	NotificationCancelFailureOrg `xml:"notificationCancelFailureOrg"` /* Организация, осуществляющая размещение отмены извещения определения поставщика (подрядчика, исполнителя) */
	CancelReason                 zfcs_protocolModificationReasonType  `xml:"cancelReason"` /* Основание для отмены извещений об отмене определения поставщика (подрядчика, исполнителя) */
	AddInfo                      zfcs_longTextType                    `xml:"addInfo"`      /* Дополнительная информация */
	Attachments                  zfcs_attachmentListType              `xml:"attachments"`  /* Прикрепленные документы */
}

type zfcs_notificationEFType struct {
	BasedOn       zfcs_purchaseNotificationType
	ETP           zfcs_ETPType `xml:"ETP"` /* Электронная торговая площадка */
	ProcedureInfo struct {
		Collecting zfcs_purchaseProcedureCollectingType `xml:"collecting"` /* Информация о подаче заявок */
		Scoring    `xml:"scoring"`                      /* Информация о процедуре рассмотрения и оценки заявок на участие в аукционе */
		Bidding    `xml:"bidding"`                      /* Информация о процедуре проведения аукциона в электронной форме */
	} `xml:"procedureInfo"`
	Lot struct {
		MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`               /* Начальная (максимальная) цена контрактов */
		PriceFormula           zfcs_longTextType               `xml:"priceFormula"`           /* Формула цены. Устанавливается, если закупка осуществляется в соответствии с ПП РФ от 13.01.2014 №19 "Об установлении случаев, в которых при заключении контракта в документации о закупке указываются формула цены и максимальное значение цены контракта" */
		StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракта, типовых условий контракта */
		Currency               zfcs_currencyRef                `xml:"currency"`               /* Валюта */
		FinanceSource          zfcs_longTextType               `xml:"financeSource"`          /* Источник финансирования */
		QuantityUndefined      bool                            `xml:"quantityUndefined"`      /* Невозможно определить количество товара, объем подлежащих выполнению работ, оказанию услуг */
		CustomerRequirements   `xml:"customerRequirements"`    /* Требования заказчиков */
		PurchaseObjects        `xml:"purchaseObjects"`         /* Объекты закупки */
		Preferenses            `xml:"preferenses"`             /* Преимущества */
		Requirements           `xml:"requirements"`            /* Требования */
		RestrictInfo           zfcs_longTextType               `xml:"restrictInfo"`         /* Ограничения участия в определении поставщика (подрядчика, исполнителя) */
		RestrictForeignsInfo   zfcs_longTextType               `xml:"restrictForeignsInfo"` /* Условия, запреты и ограничения допуска товаров, происходящих из иностранного государства или группы иностранных государств, работ, услуг, соответственно выполняемых, оказываемых иностранными лицами (согласно п.7 ч.5 ст.63 Федерального закона № 44-ФЗ) */
		AddInfo                zfcs_longTextType               `xml:"addInfo"`              /* Дополнительная информация */
		NoPublicDiscussion     bool                            `xml:"noPublicDiscussion"`   /* Закупка не подлежит обязательному общественному обсуждению в соответствии с подпунктами 2) и 3) пункта 1.4 Приказа Минэкономразвития от 10.10.2013 г. № 578  */
		PublicDiscussion       `xml:"publicDiscussion"`        /* Общественное обсуждение крупных закупок */
	} `xml:"lot"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationEFDateChangeType struct {
	Id                  int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId          zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber      zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocNumber           zfcs_documentNumberType `xml:"docNumber"`      /* Номер уведомления об изменении даты и времени проведения ЭА */
	DocPublishDate      string                  `xml:"docPublishDate"` /* Дата публикации документа */
	PurchaseResponsible struct {
		ResponsibleOrg  zfcs_organizationRef     `xml:"responsibleOrg"`
		ResponsibleRole zfcs_responsibleRoleType `xml:"responsibleRole"`
	} `xml:"purchaseResponsible"`
	PurchaseObjectInfo zfcs_longTextType  `xml:"purchaseObjectInfo"` /* Наименование объекта закупки (для печатной формы) */
	PlacingWay         `xml:"placingWay"` /* Подспособ определения поставщика (для печатной формы) */
	AuctionTime        string             `xml:"auctionTime"`    /* Дата и время проведения электронного аукциона в действующем извещении */
	NewAuctionDate     string             `xml:"newAuctionDate"` /* Новая дата проведения ЭА */
	Reason             struct {
		AddInfo zfcs_longTextType `xml:"addInfo"`
	} `xml:"reason"`
	Href         zfcs_hrefType         `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType    `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_notificationEPType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	DocNumber           zfcs_documentNumberType `xml:"docNumber"`          /* Номер документа */
	Href                zfcs_hrefType           `xml:"href"`               /* Гиперссылка на опубликованный документ */
	PrintForm           zfcs_printFormType      `xml:"printForm"`          /* Печатная форма документа */
	ExtPrintForm        zfcs_extPrintFormType   `xml:"extPrintForm"`       /* Электронный документ, полученный из внешней системы */
	PurchaseObjectInfo  zfcs_longTextType       `xml:"purchaseObjectInfo"` /* Наименование объекта закупки */
	PurchaseResponsible struct {
		ResponsibleOrg  zfcs_purchaseOrganizationType `xml:"responsibleOrg"`
		ResponsibleRole zfcs_responsibleRoleType      `xml:"responsibleRole"`
		ResponsibleInfo zfcs_contactInfoType          `xml:"responsibleInfo"`
		SpecializedOrg  zfcs_purchaseOrganizationType `xml:"specializedOrg"`
	} `xml:"purchaseResponsible"`
	PlacingWay zfcs_placingWayType `xml:"placingWay"` /* Подспособ определения поставщика */
	Lot        struct {
		MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`
		PriceFormula           zfcs_longTextType               `xml:"priceFormula"`
		StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"`
		Currency               zfcs_currencyRef                `xml:"currency"`
		FinanceSource          zfcs_longTextType               `xml:"financeSource"`
		QuantityUndefined      bool                            `xml:"quantityUndefined"`
		CustomerRequirements   `xml:"customerRequirements"`
		PurchaseObjects        `xml:"purchaseObjects"`
		Preferenses            `xml:"preferenses"`
		Requirements           `xml:"requirements"`
		RestrictInfo           zfcs_longTextType `xml:"restrictInfo"`
		NoPublicDiscussion     bool              `xml:"noPublicDiscussion"`
		PublicDiscussion       `xml:"publicDiscussion"`
	} `xml:"lot"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
	Okpd2okved2  bool                              `xml:"okpd2okved2"`  /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
}

type zfcs_notificationI111Type struct {
	BasedOn       zfcs_purchaseNotificationISType
	ProcedureInfo struct {
		CollectingEndDate string `xml:"collectingEndDate"` /* Дата и время окончания подачи заявок */
	} `xml:"procedureInfo"`
	Lots struct {
		Lot zfcs_lotI111Type `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments               zfcs_attachmentListType           `xml:"attachments"`               /* Документация об аукционе */
	Modification              zfcs_notificationModificationType `xml:"modification"`              /* Основание внесения изменений */
	ParticularsActProcurement string                            `xml:"particularsActProcurement"` /* Реквизиты нормативно-правового акта, являющегося основанием для осуществления закупки с учетом положений статьи 111 Федерального закона № 44-ФЗ */
}

type zfcs_notificationISMType struct {
	BasedOn       zfcs_purchaseNotificationISType
	ProcedureInfo struct {
		CollectingEndDate string `xml:"collectingEndDate"` /* Дата и время окончания подачи заявок */
	} `xml:"procedureInfo"`
	Lots struct {
		Lot zfcs_lotISType `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationISOType struct {
	BasedOn       zfcs_purchaseNotificationISType
	ProcedureInfo struct {
		CollectingEndDate string `xml:"collectingEndDate"` /* Дата и время окончания подачи заявок */
	} `xml:"procedureInfo"`
	Lot          zfcs_lotISType                    `xml:"lot"`          /* Лот извещения */
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationLotCancelType struct {
	BasedOn zfcs_purchaseDocumentCommonType
	Lot     struct {
		LotNumber uint64 `xml:"lotNumber"` /* Номер лота в извещении */
	} `xml:"lot"`
	CancelReason zfcs_purchaseCancelReasonType `xml:"cancelReason"` /* Причина отмены определения поставщика */
	AddInfo      string                        `xml:"addInfo"`      /* Дополнительная информация */
}

type zfcs_notificationLotChangeType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	Lot                   `xml:"lot"`                           /* Лот извещения */
	Attachments           zfcs_attachmentListType               `xml:"attachments"`  /* Документация об аукционе */
	Modification          zfcs_notificationModificationType     `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationModificationType struct {
	ModificationNumber int                     `xml:"modificationNumber"` /* Номер изменения */
	Info               zfcs_longTextType       `xml:"info"`               /* Краткое описание изменения */
	AddInfo            zfcs_longTextType       `xml:"addInfo"`            /* Дополнительная информация */
	Reason             zfcs_purchaseChangeType `xml:"reason"`             /* Основание внесения исправлений */
}

type zfcs_notificationOrgChangeType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	DocNumber      zfcs_documentNumberType `xml:"docNumber"`      /* Номер документа */
	DocDate        string                  `xml:"docDate"`        /* Дата документа */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	 */
	BaseChange struct {
		ChangeType zfcs_notificationChangeEnumType `xml:"changeType"`
		AddInfo    zfcs_longTextType               `xml:"addInfo"`
	} `xml:"baseChange"`
	NotifRespOrg `xml:"notifRespOrg"` /* Организация, осуществляющая размещение уведомления */
	Purchase     struct {
		PurchaseNumber     zfcs_purchaseNumberType `xml:"purchaseNumber"`
		PurchaseObjectInfo zfcs_longTextType       `xml:"purchaseObjectInfo"`
	} `xml:"purchase"`
	PreviousRespOrg `xml:"previousRespOrg"` /* Организация, осуществлявшая закупку */
	NewRespOrg      `xml:"newRespOrg"`      /* Сведения о новой организации, осуществляющей закупку */
	Href            zfcs_hrefType           `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm       zfcs_printFormType      `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm    zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	Attachments     zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
}

type zfcs_notificationOKDType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	ProcedureInfo         zfcs_purchaseProcedureOKDType         `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationOKOUType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	ProcedureInfo         zfcs_purchaseProcedureOKOUType        `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationOKType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	ProcedureInfo         zfcs_purchaseProcedureOKType          `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationPOType struct {
	BasedOn             zfcs_purchaseNotificationType
	ContractServiceInfo zfcs_longTextType `xml:"contractServiceInfo"` /* Информация о контрактной службе, контрактном управляющем */
	ProcedureInfo       struct {
		Collecting  zfcs_purchaseProcedureCollectingWithFormType `xml:"collecting"`  /* Информация о подаче заявок на участие в предварительном отборе  */
		Selecting   zfcs_purchaseProcedureSelectingType          `xml:"selecting"`   /* Информация о проведении предварительного отбора */
		Contracting zfcs_purchaseProcedureContractingType        `xml:"contracting"` /* Информация о заключении контракта */
	} `xml:"procedureInfo"`
	Lot struct {
		MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`               /* Начальная (максимальная) цена контрактов */
		PriceFormula           zfcs_longTextType               `xml:"priceFormula"`           /* Формула цены. Устанавливается, если закупка осуществляется в соответствии с ПП РФ от 13.01.2014 №19 "Об установлении случаев, в которых при заключении контракта в документации о закупке указываются формула цены и максимальное значение цены контракта" */
		StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракта, типовых условий контракта */
		Currency               zfcs_currencyRef                `xml:"currency"`               /* Валюта */
		FinanceSource          zfcs_longTextType               `xml:"financeSource"`          /* Источник финансирования */
		BeforePay              zfcs_longTextType               `xml:"beforePay"`              /* Информация о необходимости без предварительной оплаты и (или) с отсрочкой платежа осуществить поставки товаров, выполнение работ, оказание услуг в возможно короткий срок */
		QuantityUndefined      bool                            `xml:"quantityUndefined"`      /* Невозможно определить количество товара, объем подлежащих выполнению работ, оказанию услуг */
		CustomerRequirements   `xml:"customerRequirements"`    /* Требования заказчиков */
		PurchaseObjects        `xml:"purchaseObjects"`         /* Объекты закупки */
		Preferenses            `xml:"preferenses"`             /* Преимущества */
		Requirements           `xml:"requirements"`            /* Требования */
		RestrictInfo           zfcs_longTextType               `xml:"restrictInfo"` /* Ограничение участия в определении поставщика (подрядчика, исполнителя), установленное в соответствии с ФЗ (согласно п.4 ст.42 Федерального закона № 44-ФЗ) */
		AddInfo                zfcs_longTextType               `xml:"addInfo"`      /* Дополнительная информация */
	} `xml:"lot"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationZakAType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении аукционной документации */
	ProcedureInfo         zfcs_purchaseProcedureZakAType        `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot zfcs_lotOKType `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationZakKDType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	ProcedureInfo         zfcs_purchaseProcedureOKDType         `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot zfcs_lotOKType `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация о конкурсе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationZakKOUType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	ProcedureInfo         zfcs_purchaseProcedureOKOUType        `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot zfcs_lotOKType `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация о конкурсе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationZakKType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении конкурсной документации */
	ProcedureInfo         zfcs_purchaseProcedureOKType          `xml:"procedureInfo"`         /* Информация о процедуре закупки */
	Lots                  struct {
		Lot zfcs_lotOKType `xml:"lot"` /* Лот извещения */
	} `xml:"lots"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация о конкурсе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationZKType struct {
	BasedOn             zfcs_purchaseNotificationType
	ContractServiceInfo zfcs_longTextType `xml:"contractServiceInfo"` /* Информация о контрактной службе, контрактном управляющем */
	ProcedureInfo       struct {
		Collecting  zfcs_purchaseProcedureCollectingWithFormType `xml:"collecting"`  /* Информация о подаче заявок */
		Opening     zfcs_purchaseProcedureOpeningType            `xml:"opening"`     /* Информация о вскрытии конвертов, открытии доступа к электронным документам заявок участников */
		Contracting zfcs_purchaseProcedureContractingType        `xml:"contracting"` /* Информация о заключении контракта */
	} `xml:"procedureInfo"`
	Lot          `xml:"lot"`                       /* Лот извещения */
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_notificationZPType struct {
	BasedOn               zfcs_purchaseNotificationType
	PurchaseDocumentation zfcs_releasePurchaseDocumentationType `xml:"purchaseDocumentation"` /* Информация о предоставлении документации о проведении запроса предложений */
	ProcedureInfo         struct {
		Collecting   zfcs_purchaseProcedureCollectingType   `xml:"collecting"`   /* Информация о подаче заявок */
		Opening      zfcs_purchaseProcedureOpeningType      `xml:"opening"`      /* Информация о вскрытии конвертов, открытии доступа к электронным документам заявок участников */
		Scoring      zfcs_purchaseProcedureScoringPlaceType `xml:"scoring"`      /* Информация о процедуре рассмотрения и оценки заявок на участие в конкурсе */
		FinalOpening zfcs_purchaseProcedureOpeningType      `xml:"finalOpening"` /* Информация о вскрытии конвертов с окончательными предложениями, открытии доступа к электронным документам окончательных предложений участников */
	} `xml:"procedureInfo"`
	Lot struct {
		MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`               /* Начальная (максимальная) цена контрактов */
		PriceFormula           zfcs_longTextType               `xml:"priceFormula"`           /* Формула цены. Устанавливается, если закупка осуществляется в соответствии с ПП РФ от 13.01.2014 №19 "Об установлении случаев, в которых при заключении контракта в документации о закупке указываются формула цены и максимальное значение цены контракта" */
		StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракта, типовых условий контракта */
		Currency               zfcs_currencyRef                `xml:"currency"`               /* Валюта */
		FinanceSource          zfcs_longTextType               `xml:"financeSource"`          /* Источник финансирования */
		QuantityUndefined      bool                            `xml:"quantityUndefined"`      /* Невозможно определить количество товара, объем подлежащих выполнению работ, оказанию услуг */
		CustomerRequirements   `xml:"customerRequirements"`    /* Требования заказчиков */
		PurchaseObjects        `xml:"purchaseObjects"`         /* Объекты закупки */
		Preferenses            `xml:"preferenses"`             /* Преимущества */
		Requirements           `xml:"requirements"`            /* Требования */
		RestrictInfo           zfcs_longTextType               `xml:"restrictInfo"`       /* Ограничение участия в определении поставщика (подрядчика, исполнителя), установленное в соответствии с ФЗ (согласно п.4 ст.42 Федерального закона № 44-ФЗ) */
		AddInfo                zfcs_longTextType               `xml:"addInfo"`            /* Дополнительная информация */
		PublicDiscussion       zfcs_publicDiscussionType       `xml:"publicDiscussion"`   /* Общественное обсуждение */
		NoPublicDiscussion     bool                            `xml:"noPublicDiscussion"` /* Закупка не подлежит обязательному общественному обсуждению в соответствии с подпунктами 2) и 3) пункта 1.4 Приказа Минэкономразвития от 10.10.2013 г. № 578  */
	} `xml:"lot"`
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация об аукционе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_timeEFType struct {
	PurchaseNumber                 zfcs_purchaseNumberType `xml:"purchaseNumber"`                 /* Номер закупки */
	DocNumber                      zfcs_documentNumberType `xml:"docNumber"`                      /* Номер уведомления об изменении даты и времени проведения ЭА */
	AuctionTime                    string                  `xml:"auctionTime"`                    /* Время проведения */
	NotificationModificationNumber int                     `xml:"notificationModificationNumber"` /* Номер изменения извещения по закупке, для которого задается время проведения аукциона */
}

type zfcs_docCancelReasonType struct {
}

type zfcs_budgetFinancingsType struct {
	BudgetFinancing struct {
		Year zfcs_yearType  `xml:"year"`
		Sum  zfcs_moneyType `xml:"sum"`
	} `xml:"budgetFinancing"`
	TotalSum zfcs_moneyType `xml:"totalSum"` /* Общая сумма бюджетного финансирования */
}

type zfcs_nonbudgetFinancingsType struct {
	NonbudgetFinancing struct {
		Year zfcs_yearType  `xml:"year"`
		Sum  zfcs_moneyType `xml:"sum"`
	} `xml:"nonbudgetFinancing"`
	TotalSum zfcs_moneyType `xml:"totalSum"` /* Общая сумма внебюджетного финансирования */
}

type zfcs_ETPType struct {
	Code string `xml:"code"` /* Кодовое наименование ЭП */
	Name string `xml:"name"` /* Наименование ЭТП */
	Url  string `xml:"url"`  /* Адрес ЭП */
}

type zfcs_guaranteeType struct {
	Amount            zfcs_moneyType             `xml:"amount"`            /* Размер обеспечения */
	Part              float64                    `xml:"part"`              /* Доля от начальной (максимальной) цены контракта */
	ProcedureInfo     string                     `xml:"procedureInfo"`     /* Порядок внесения денежных средств в качестве обеспечения заявки (порядок предоставления обеспечения исполнения контракта) */
	SettlementAccount zfcs_settlementAccountType `xml:"settlementAccount"` /* Номер расчётного счёта */
	PersonalAccount   zfcs_personalAccountType   `xml:"personalAccount"`   /* Номер лицевого счёта  */
	Bik               zfcs_bikType               `xml:"bik"`               /* БИК */
}

type zfcs_kladrPlacesType struct {
	KladrPlace struct {
		DeliveryPlace              zfcs_longTextType `xml:"deliveryPlace"`
		NoKladrForRegionSettlement `xml:"noKladrForRegionSettlement"`
	} `xml:"kladrPlace"`
}

type zfcs_kladrType struct {
	KladrType string `xml:"kladrType"` /* Тип элемента КЛАДР */
	KladrCode string `xml:"kladrCode"` /* Код КЛАДР */
	FullName  string `xml:"fullName"`  /* Полное наименование */
}

type zfcs_lotI111Type struct {
	LotNumber   uint64           `xml:"lotNumber"` /* Номер лота в извещении */
	MaxPrice    zfcs_moneyType   `xml:"maxPrice"`  /* Начальная (максимальная) цена контрактов. Отсутствие элемента в извещении означает, что НМЦК не установлена. */
	Currency    zfcs_currencyRef `xml:"currency"`  /* Валюта */
	OKPD2       zfcs_OKPDRef     `xml:"OKPD2"`     /* Классификация по ОКПД2 */
	Preferenses struct {
		Preferense zfcs_preferenseType `xml:"preferense"`
	} `xml:"preferenses"`
	Requirements struct {
		Requirement zfcs_requirementType `xml:"requirement"`
	} `xml:"requirements"`
	AddInfo                zfcs_longTextType         `xml:"addInfo"`                /* Дополнительная информация */
	PublicDiscussion       zfcs_publicDiscussionType `xml:"publicDiscussion"`       /* Информация об общественных слушаниях по лоту закупки (для печатной формы) */
	MaxCostDefinitionOrder zfcs_longTextType         `xml:"maxCostDefinitionOrder"` /* Порядок определения начальной (максимальной) цены контракта в случаях, установленных Правительством Российской Федерации */
}

type zfcs_lotISType struct {
	LotNumber   uint64           `xml:"lotNumber"` /* Номер лота в извещении */
	MaxPrice    zfcs_moneyType   `xml:"maxPrice"`  /* Начальная (максимальная) цена контрактов. Отсутствие элемента в извещении означает, что НМЦК не установлена. */
	Currency    zfcs_currencyRef `xml:"currency"`  /* Валюта */
	Preferenses struct {
		Preferense zfcs_preferenseType `xml:"preferense"`
	} `xml:"preferenses"`
	Requirements struct {
		Requirement zfcs_requirementType `xml:"requirement"`
	} `xml:"requirements"`
	AddInfo          zfcs_longTextType         `xml:"addInfo"`          /* Дополнительная информация */
	PublicDiscussion zfcs_publicDiscussionType `xml:"publicDiscussion"` /* Информация об общественных слушаниях по лоту закупки (для печатной формы) */
}

type zfcs_lotOKType struct {
	LotNumber              uint64                          `xml:"lotNumber"`              /* Номер лота в извещении */
	LotObjectInfo          zfcs_longTextType               `xml:"lotObjectInfo"`          /* Наименование объекта закупки для лота */
	MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`               /* Начальная (максимальная) цена контрактов */
	PriceFormula           zfcs_longTextType               `xml:"priceFormula"`           /* Формула цены. Устанавливается, если закупка осуществляется в соответствии с ПП РФ от 13.01.2014 №19 "Об установлении случаев, в которых при заключении контракта в документации о закупке указываются формула цены и максимальное значение цены контракта" */
	StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракта, типовых условий контракта */
	Currency               zfcs_currencyRef                `xml:"currency"`               /* Валюта */
	FinanceSource          zfcs_longTextType               `xml:"financeSource"`          /* Источник финансирования */
	QuantityUndefined      bool                            `xml:"quantityUndefined"`      /* Невозможно определить количество товара, объем подлежащих выполнению работ, оказанию услуг */
	CustomerRequirements   struct {
		CustomerRequirement `xml:"customerRequirement"`
	} `xml:"customerRequirements"`
	PurchaseObjects struct {
		PurchaseObject `xml:"purchaseObject"`
		TotalSum       zfcs_moneyType `xml:"totalSum"`
	} `xml:"purchaseObjects"`
	Preferenses struct {
		Preferense zfcs_preferenseType `xml:"preferense"`
	} `xml:"preferenses"`
	Requirements struct {
		Requirement zfcs_requirementType `xml:"requirement"`
	} `xml:"requirements"`
	RestrictInfo         zfcs_longTextType         `xml:"restrictInfo"`         /* Ограничение участия в определении поставщика (подрядчика, исполнителя), установленное в соответствии с ФЗ (согласно п.4 ст.42 Федерального закона № 44-ФЗ) */
	RestrictForeignsInfo zfcs_longTextType         `xml:"restrictForeignsInfo"` /* Условия, запреты и ограничения допуска товаров, происходящих из иностранного государства или группы иностранных государств, работ, услуг, соответственно выполняемых, оказываемых иностранными лицами (согласно п.8 ч.3 ст.49 Федерального закона № 44-ФЗ) */
	AddInfo              zfcs_longTextType         `xml:"addInfo"`              /* Дополнительная информация */
	PublicDiscussion     zfcs_publicDiscussionType `xml:"publicDiscussion"`     /* Общественное обсуждение крупных закупок */
}

type zfcs_lotZKType struct {
	MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`               /* Начальная (максимальная) цена контрактов */
	MaxPriceInfo           zfcs_longTextType               `xml:"maxPriceInfo"`           /* Обоснование начальной (максимальной) цены контракта  */
	PriceFormula           zfcs_longTextType               `xml:"priceFormula"`           /* Формула цены. Устанавливается, если закупка осуществляется в соответствии с ПП РФ от 13.01.2014 №19 "Об установлении случаев, в которых при заключении контракта в документации о закупке указываются формула цены и максимальное значение цены контракта" */
	StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракта, типовых условий контракта */
	Currency               zfcs_currencyRef                `xml:"currency"`               /* Валюта */
	FinanceSource          zfcs_longTextType               `xml:"financeSource"`          /* Источник финансирования */
	QuantityUndefined      bool                            `xml:"quantityUndefined"`      /* Невозможно определить количество товара, объем подлежащих выполнению работ, оказанию услуг */
	CustomerRequirements   struct {
		CustomerRequirement `xml:"customerRequirement"`
	} `xml:"customerRequirements"`
	PurchaseObjects struct {
		PurchaseObject `xml:"purchaseObject"`
		TotalSum       zfcs_moneyType `xml:"totalSum"`
	} `xml:"purchaseObjects"`
	Preferenses struct {
		Preferense zfcs_preferenseType `xml:"preferense"`
	} `xml:"preferenses"`
	Requirements struct {
		Requirement zfcs_requirementType `xml:"requirement"`
	} `xml:"requirements"`
	RestrictInfo     zfcs_longTextType         `xml:"restrictInfo"`     /* Ограничение участия в определении поставщика (подрядчика, исполнителя), установленное в соответствии с ФЗ (согласно п.4 ст.42 Федерального закона № 44-ФЗ) */
	AddInfo          zfcs_longTextType         `xml:"addInfo"`          /* Дополнительная информация */
	PublicDiscussion zfcs_publicDiscussionType `xml:"publicDiscussion"` /* Информация об общественных слушаниях по лоту закупки (для печатной формы) */
}

type zfcs_lotInfoType struct {
	LotObjectInfo          zfcs_longTextType               `xml:"lotObjectInfo"`          /* Наименование объекта закупки для лота */
	Currency               zfcs_currencyRef                `xml:"currency"`               /* Валюта */
	MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`               /* Начальная (максимальная) цена контрактов или цена за единицу измерения */
	SpelledMaxPrice        zfcs_longTextType               `xml:"spelledMaxPrice"`        /* Начальная (максимальная) цена контракта (НМЦК) прописью */
	UnitPrice              zfcs_moneyType                  `xml:"unitPrice"`              /* Цена за единицу измерения */
	SpelledUnitPrice       zfcs_longTextType               `xml:"spelledUnitPrice"`       /* Цена за единицу измерения прописью */
	StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"` /* Номер типового контракта, типовых условий контракта */
	PriceFormula           zfcs_longTextType               `xml:"priceFormula"`           /* Формула цены. Устанавливается, если закупка осуществляется в соответствии с ПП РФ от 13.01.2014 №19 "Об установлении случаев, в которых при заключении контракта в документации о закупке указываются формула цены и максимальное значение цены контракта" */
	FinanceSource          zfcs_longTextType               `xml:"financeSource"`          /* Источник финансирования */
	DeliveryTerm           zfcs_longTextType               `xml:"deliveryTerm"`           /* Сроки доставки товара, выполнения работы или оказания услуги либо график оказания услуг */
	Customers              struct {
		Customer zfcs_organizationRef `xml:"customer"`
	} `xml:"customers"`
	Preferenses struct {
		Preferense zfcs_preferenseType `xml:"preferense"`
	} `xml:"preferenses"`
	Requirements struct {
		Requirement zfcs_requirementType `xml:"requirement"`
	} `xml:"requirements"`
	PurchaseProlongation struct {
		ProlongationNumber zfcs_documentNumberType `xml:"prolongationNumber"`
		PublishDate        string                  `xml:"publishDate"`
	} `xml:"purchaseProlongation"`
	SpelledInvalidAppCount zfcs_longTextType `xml:"spelledInvalidAppCount"` /* Количество заявок, не соответствующих установленным единым и дополнительным требованиям прописью */
	SpelledAppCount        zfcs_longTextType `xml:"spelledAppCount"`        /* Количество поданных заявок по лоту из связанного протокола */
	SpelledValidAppCount   zfcs_longTextType `xml:"spelledValidAppCount"`   /* Количество заявок, соответствующих установленным единым и дополнительным требованиям прописью */
}

type zfcs_paymentInfoType struct {
	Amount            zfcs_moneyType             `xml:"amount"`            /* Размер обеспечения */
	Part              float64                    `xml:"part"`              /* Доля от начальной (максимальной) цены контракта */
	ProcedureInfo     zfcs_longTextType          `xml:"procedureInfo"`     /* Порядок внесения денежных средств в качестве обеспечения заявки (порядок предоставления обеспечения исполнения контракта) */
	SettlementAccount zfcs_settlementAccountType `xml:"settlementAccount"` /* Номер расчётного счёта */
	PersonalAccount   zfcs_personalAccountType   `xml:"personalAccount"`   /* Номер лицевого счёта  */
	Bik               zfcs_bikType               `xml:"bik"`               /* БИК */
}

type zfcs_placingWayType struct {
	Code zfcs_nsiPlacingWayCodeType `xml:"code"` /* Код подспособа определения поставщика */
	Name string                     `xml:"name"` /* Наименование подспособа определения поставщика */
}

type zfcs_preferenseType struct {
	Code      int64   `xml:"code"`      /* Код преимущества */
	Name      string  `xml:"name"`      /* Наименование премущества */
	PrefValue float64 `xml:"prefValue"` /* Величина (преимущества) */
}

type zfcs_publicDiscussionType struct {
	Place zfcs_publicDiscussionPlaceEnum `xml:"place"` /* Место проведения общественного обсуждения:
	E - в разделе «Общественные обсуждения крупных закупок» Официального сайта Российской Федерации в сети Интернет для размещения информации о размещении заказов на поставки товаров, выполнение работ, оказание услуг;
	F - на форуме Официального сайта Российской Федерации в сети Интернет. */
}

type zfcs_purchaseApprovalType struct {
	ReceiptNumber  string               `xml:"receiptNumber"`  /* Номер квитанции */
	Authority      zfcs_organizationRef `xml:"authority"`      /* Уполномоченный орган (учреждение) */
	ApprovalDate   string               `xml:"approvalDate"`   /* Дата принятия решения  */
	ApprovalResult bool                 `xml:"approvalResult"` /* Согласование публикации извещения (изменения извещения) */
	Reason         zfcs_longTextType    `xml:"reason"`         /* Причина отказа */
	AddInfo        zfcs_longTextType    `xml:"addInfo"`        /* Дополнительная информация */
	PublishDate    string               `xml:"publishDate"`    /* Дата публикации */
	Purchase       struct {
		PurchaseNumber     zfcs_purchaseNumberType `xml:"purchaseNumber"`
		DocNumber          zfcs_documentNumberType `xml:"docNumber"`
		Href               zfcs_hrefType           `xml:"href"`
		ResponsibleOrg     zfcs_organizationRef    `xml:"responsibleOrg"`
		PurchaseObjectInfo zfcs_longTextType       `xml:"purchaseObjectInfo"`
		PlacingWay         zfcs_placingWayType     `xml:"placingWay"`
	} `xml:"purchase"`
	Documents struct {
		DocumentInfo `xml:"documentInfo"`
	} `xml:"documents"`
}

type zfcs_purchaseCancelReasonType struct {
}

type zfcs_purchaseChangeType struct {
}

type zfcs_purchaseDocumentCancelType struct {
	BasedOn      zfcs_purchaseDocumentCommonType
	AddInfo      zfcs_longTextType        `xml:"addInfo"`      /* Дополнительная информация */
	Attachments  zfcs_attachmentListType  `xml:"attachments"`  /* Информация о прикрепленных документах */
	CancelReason zfcs_docCancelReasonType `xml:"cancelReason"` /* Причина отмены документа */
}

type zfcs_purchaseDocumentCommonType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocNumber      zfcs_documentNumberType `xml:"docNumber"`      /* Номер документа */
	DocDate        string                  `xml:"docDate"`        /* Дата документа */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	Href         zfcs_hrefType         `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType    `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_purchaseDocumentType struct {
	BasedOn      zfcs_purchaseDocumentCommonType
	DocType      zfcs_docType            `xml:"docType"`     /* Тип документа */
	AddInfo      zfcs_longTextType       `xml:"addInfo"`     /* Дополнительная информация */
	Attachments  zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах */
	Modification struct {
		ModificationNumber int                     `xml:"modificationNumber"` /* Номер изменения */
		Info               zfcs_longTextType       `xml:"info"`               /* Краткое описание изменения */
		AddInfo            string                  `xml:"addInfo"`            /* Дополнительная информация */
		Reason             zfcs_purchaseChangeType `xml:"reason"`             /* Основание внесения изменений */
	} `xml:"modification"`
}

type zfcs_purchaseNotificationISType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	DocNumber           zfcs_documentNumberType `xml:"docNumber"`          /* Номер документа */
	Href                zfcs_hrefType           `xml:"href"`               /* Гиперссылка на опубликованный документ */
	PrintForm           zfcs_printFormType      `xml:"printForm"`          /* Печатная форма документа */
	PurchaseObjectInfo  zfcs_longTextType       `xml:"purchaseObjectInfo"` /* Наименование объекта закупки */
	PurchaseResponsible struct {
		ResponsibleOrg  zfcs_purchaseOrganizationType `xml:"responsibleOrg"`
		ResponsibleRole zfcs_responsibleRoleType      `xml:"responsibleRole"`
	} `xml:"purchaseResponsible"`
	PlacingWay  zfcs_placingWayType `xml:"placingWay"`  /* Подспособ определения поставщика */
	Okpd2okved2 bool                `xml:"okpd2okved2"` /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
}

type zfcs_purchaseNotificationType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	DocNumber           zfcs_documentNumberType `xml:"docNumber"`          /* Номер документа */
	Href                zfcs_hrefType           `xml:"href"`               /* Гиперссылка на опубликованный документ */
	PrintForm           zfcs_printFormType      `xml:"printForm"`          /* Печатная форма документа */
	ExtPrintForm        zfcs_extPrintFormType   `xml:"extPrintForm"`       /* Электронный документ, полученный из внешней системы */
	PurchaseObjectInfo  zfcs_longTextType       `xml:"purchaseObjectInfo"` /* Наименование объекта закупки */
	PurchaseResponsible struct {
		ResponsibleOrg     zfcs_purchaseOrganizationType `xml:"responsibleOrg"`
		ResponsibleRole    zfcs_responsibleRoleType      `xml:"responsibleRole"`
		ResponsibleInfo    zfcs_contactInfoType          `xml:"responsibleInfo"`
		SpecializedOrg     zfcs_purchaseOrganizationType `xml:"specializedOrg"`
		LastSpecializedOrg zfcs_purchaseOrganizationType `xml:"lastSpecializedOrg"`
	} `xml:"purchaseResponsible"`
	PlacingWay  zfcs_placingWayType `xml:"placingWay"`  /* Подспособ определения поставщика */
	Okpd2okved2 bool                `xml:"okpd2okved2"` /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
}

type zfcs_purchaseOrganizationType struct {
	RegNum          zfcs_string              `xml:"regNum"`          /* Код по СПЗ */
	ConsRegistryNum zfcs_consRegistryNumType `xml:"consRegistryNum"` /* Код по Сводному Реестру */
	FullName        zfcs_longTextType        `xml:"fullName"`        /* Полное наименование */
	ShortName       zfcs_longTextType        `xml:"shortName"`       /* Сокращенное наименование */
	PostAddress     zfcs_longTextType        `xml:"postAddress"`     /* Почтовый адрес организации */
	FactAddress     zfcs_longTextType        `xml:"factAddress"`     /* Адрес местонахождения организации */
	INN             zfcs_innOrganizationType `xml:"INN"`             /* ИНН организации */
	KPP             zfcs_kppType             `xml:"KPP"`             /* КПП организации */
}

type zfcs_purchaseProcedureBiddingType struct {
	Date    string            `xml:"date"`    /* Дата и время проведения закрытого аукциона */
	Place   zfcs_longTextType `xml:"place"`   /* Место проведения закрытого аукциона */
	AddInfo zfcs_longTextType `xml:"addInfo"` /* Дополнительная информация */
}

type zfcs_purchaseProcedureCollectingType struct {
	StartDate string            `xml:"startDate"` /* Дата и время начала подачи заявок */
	Place     zfcs_longTextType `xml:"place"`     /* Место подачи заявок */
	Order     zfcs_longTextType `xml:"order"`     /* Порядок подачи заявок */
	EndDate   string            `xml:"endDate"`   /* Дата и время окончания подачи заявок */
}

type zfcs_purchaseProcedureCollectingWithFormType struct {
	StartDate string            `xml:"startDate"` /* Дата и время начала подачи заявок */
	Place     zfcs_longTextType `xml:"place"`     /* Место подачи заявок */
	Order     zfcs_longTextType `xml:"order"`     /* Порядок подачи заявок */
	EndDate   string            `xml:"endDate"`   /* Дата и время окончания подачи заявок */
	Form      zfcs_longTextType `xml:"form"`      /* Форма котировочной заявки */
}

type zfcs_purchaseProcedureContractingType struct {
	ContractingTerm zfcs_longTextType `xml:"contractingTerm"` /* Срок, в течение которого победитель запроса котировок или иной участник запроса котировок, с которым заключается контракт при уклонении победителя от заключения контракта, должен подписать контракт  */
	EvadeConditions zfcs_longTextType `xml:"evadeConditions"` /* Условия признания победителя запроса котировок или иного участника запроса котировок уклонившимися от заключения контракта  */
}

type zfcs_purchaseProcedureOKDType struct {
	StageOne struct {
		Collecting       zfcs_purchaseProcedureCollectingType       `xml:"collecting"`
		Opening          zfcs_purchaseProcedureOpeningType          `xml:"opening"`
		Scoring          zfcs_purchaseProcedureScoringPlaceType     `xml:"scoring"`
		Prequalification zfcs_purchaseProcedurePrequalificationType `xml:"prequalification"`
	} `xml:"stageOne"`
	StageTwo struct {
		Collecting zfcs_purchaseProcedureCollectingType `xml:"collecting"`
		Opening    zfcs_purchaseProcedureOpeningType    `xml:"opening"`
		Scoring    zfcs_purchaseProcedureScoringType    `xml:"scoring"`
	} `xml:"stageTwo"`
}

type zfcs_purchaseProcedureOKOUType struct {
	Collecting       zfcs_purchaseProcedureCollectingType       `xml:"collecting"`       /* Информация о подаче заявок */
	Opening          zfcs_purchaseProcedureOpeningType          `xml:"opening"`          /* Информация о вскрытии конвертов, открытии доступа к электронным документам заявок участников */
	Prequalification zfcs_purchaseProcedurePrequalificationType `xml:"prequalification"` /* Информация о предквалификационном отборе */
	Scoring          zfcs_purchaseProcedureScoringType          `xml:"scoring"`          /* Информация о процедуре рассмотрения и оценки заявок на участие в конкурсе */
}

type zfcs_purchaseProcedureOKType struct {
	Collecting zfcs_purchaseProcedureCollectingType `xml:"collecting"` /* Информация о подаче заявок */
	Opening    zfcs_purchaseProcedureOpeningType    `xml:"opening"`    /* Информация о вскрытии конвертов, открытии доступа к электронным документам заявок участников */
	Scoring    zfcs_purchaseProcedureScoringType    `xml:"scoring"`    /* Информация о процедуре рассмотрения и оценки заявок на участие в конкурсе */
}

type zfcs_purchaseProcedureOpeningType struct {
	Date    string            `xml:"date"`    /* Дата и время вскрытия конвертов, открытии доступа к электронным документам заявок участников */
	Place   zfcs_longTextType `xml:"place"`   /* Место вскрытия конвертов, открытии доступа к электронным документам заявок участников */
	AddInfo zfcs_longTextType `xml:"addInfo"` /* Дополнительная информация */
}

type zfcs_purchaseProcedurePrequalificationType struct {
	Date  string            `xml:"date"`  /* Дата и время предквалификационного отбора */
	Place zfcs_longTextType `xml:"place"` /* Место проведения предквалификационного отбора */
}

type zfcs_purchaseProcedureScoringPlaceType struct {
	Date    string            `xml:"date"`    /* Дата рассмотрения и оценки заявок на участие в конкурсе */
	Place   zfcs_longTextType `xml:"place"`   /* Место рассмотрения и оценки заявок на участие в конкурсе */
	AddInfo zfcs_longTextType `xml:"addInfo"` /* Дополнительная информация */
}

type zfcs_purchaseProcedureScoringType struct {
	Date    string            `xml:"date"`    /* Дата рассмотрения и оценки заявок на участие в конкурсе */
	Place   zfcs_longTextType `xml:"place"`   /* Место рассмотрения и оценки заявок на участие в конкурсе */
	AddInfo zfcs_longTextType `xml:"addInfo"` /* Дополнительная информация (устарело) */
}

type zfcs_purchaseProcedureSelectingType struct {
	Date  string            `xml:"date"`  /* Дата и время проведения предварительного отбора  */
	Place zfcs_longTextType `xml:"place"` /* Место проведения предварительного отбора */
}

type zfcs_purchaseProcedureZakAType struct {
	Collecting zfcs_purchaseProcedureCollectingType `xml:"collecting"` /* Информация о подаче заявок */
	Opening    zfcs_purchaseProcedureOpeningType    `xml:"opening"`    /* Информация о вскрытии конвертов, открытии доступа к электронным документам заявок участников */
	Scoring    zfcs_purchaseProcedureScoringType    `xml:"scoring"`    /* Информация о процедуре рассмотрения и оценки заявок на участие в аукционе */
	Bidding    zfcs_purchaseProcedureBiddingType    `xml:"bidding"`    /* Информация о процедуре проведения закрытого аукциона */
}

type zfcs_purchaseProlongationCommonType struct {
	BasedOn zfcs_purchaseDocumentCommonType
	DocType zfcs_docType `xml:"docType"` /* Тип документа */
}

type zfcs_purchaseProlongationOKType struct {
	BasedOn zfcs_purchaseProlongationCommonType
	Lot     struct {
		LotNumber     uint64            `xml:"lotNumber"`     /* Номер лота в извещении */
		LotObjectInfo zfcs_longTextType `xml:"lotObjectInfo"` /* Наименование объекта закупки для лота */
	} `xml:"lot"`
	ScoringDate             string `xml:"scoringDate"`             /* Дата и время рассмотрения и оценки заявок в действующей редакции извещения о закупке */
	ScoringProlongationDate string `xml:"scoringProlongationDate"` /* Продленная дата и время рассмотрения заявок */
}

type zfcs_purchaseProlongationZKType struct {
	BasedOn                    zfcs_purchaseProlongationCommonType
	CollectingEndDate          string `xml:"collectingEndDate"`          /* Дата и время окончания подачи заявок в действующей редакции извещения о закупке */
	CollectingProlongationDate string `xml:"collectingProlongationDate"` /* Новые дата и время окончания срока подачи заявок */
	OpeningDate                string `xml:"openingDate"`                /* Дата и время вскрытия конвертов, открытия доступа к электронным документам заявок участников в действующей редакции извещения о закупке */
	OpeningProlongationDate    string `xml:"openingProlongationDate"`    /* Новые дата и время вскрытия конвертов, открытия доступа к электронным документам заявок участников */
}

type zfcs_purchaseProtocolEFType struct {
	Id                       int64                   `xml:"id"`                       /* Идентификатор документа ЕИС */
	ExternalId               zfcs_externalIdType     `xml:"externalId"`               /* Внешний идентификатор документа */
	PurchaseNumber           zfcs_purchaseNumberType `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber zfcs_documentNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     zfcs_documentNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    zfcs_longTextType       `xml:"place"`                    /* Место проведения процедуры */
	ProtocolDate             string                  `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string                  `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string                  `xml:"publishDate"`              /* Дата публикации */
	Commission               zfcs_commissionType     `xml:"commission"`               /* Информация о комиссии */
	Href                     zfcs_hrefType           `xml:"href"`                     /* Гиперссылка на опубликованный документ */
	PrintForm                zfcs_printFormType      `xml:"printForm"`                /* Печатная форма протокола */
	ExtPrintForm             zfcs_extPrintFormType   `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	ProtocolPublisher        struct {
		PublisherOrg  zfcs_purchaseOrganizationType `xml:"publisherOrg"`
		PublisherRole zfcs_responsibleRoleType      `xml:"publisherRole"`
	} `xml:"protocolPublisher"`
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	Modification zfcs_protocolModificationType `xml:"modification"` /* Основание внесения исправлений */
}

type zfcs_purchaseProtocolEFNoCommissionType struct {
	Id                       int64                   `xml:"id"`                       /* Идентификатор документа ЕИС */
	ExternalId               zfcs_externalIdType     `xml:"externalId"`               /* Внешний идентификатор документа */
	PurchaseNumber           zfcs_purchaseNumberType `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber zfcs_documentNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     zfcs_documentNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    zfcs_longTextType       `xml:"place"`                    /* Место проведения процедуры */
	ProtocolDate             string                  `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string                  `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string                  `xml:"publishDate"`              /* Дата публикации */
	Commission               zfcs_commissionType     `xml:"commission"`               /* Информация о комиссии. (устарело) Игнорируется при приеме и передаче. Оставлено для совместимости с предыдущими версиями. */
	Href                     zfcs_hrefType           `xml:"href"`                     /* Гиперссылка на опубликованный документ */
	PrintForm                zfcs_printFormType      `xml:"printForm"`                /* Печатная форма протокола */
	ExtPrintForm             zfcs_extPrintFormType   `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	ProtocolPublisher        struct {
		PublisherOrg  zfcs_purchaseOrganizationType `xml:"publisherOrg"`
		PublisherRole zfcs_responsibleRoleType      `xml:"publisherRole"`
	} `xml:"protocolPublisher"`
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	Modification zfcs_protocolModificationType `xml:"modification"` /* Основание внесения исправлений */
}

type zfcs_purchaseProtocolType struct {
	Id                       int64                   `xml:"id"`                       /* Идентификатор документа ЕИС */
	ExternalId               zfcs_externalIdType     `xml:"externalId"`               /* Внешний идентификатор документа */
	PurchaseNumber           zfcs_purchaseNumberType `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber zfcs_documentNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     zfcs_documentNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    zfcs_longTextType       `xml:"place"`                    /* Место проведения процедуры */
	ProtocolDate             string                  `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string                  `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string                  `xml:"publishDate"`              /* Дата публикации */
	Commission               zfcs_commissionType     `xml:"commission"`               /* Информация о комиссии */
	Href                     zfcs_hrefType           `xml:"href"`                     /* Гиперссылка на опубликованный документ */
	PrintForm                zfcs_printFormType      `xml:"printForm"`                /* Печатная форма протокола */
	ExtPrintForm             zfcs_extPrintFormType   `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	ProtocolPublisher        struct {
		PublisherOrg  zfcs_purchaseOrganizationType `xml:"publisherOrg"`
		PublisherRole zfcs_responsibleRoleType      `xml:"publisherRole"`
	} `xml:"protocolPublisher"`
	PurchaseInfo struct {
		PurchaseResponsible    `xml:"purchaseResponsible"`
		PlacingWay             zfcs_placingWayType     `xml:"placingWay"`
		PublishDate            string                  `xml:"publishDate"`
		PurchaseObjectInfo     zfcs_longTextType       `xml:"purchaseObjectInfo"`
		NotificationFullNumber zfcs_documentNumberType `xml:"notificationFullNumber"`
		NotificationFullName   zfcs_longTextType       `xml:"notificationFullName"`
	} `xml:"purchaseInfo"`
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	Modification zfcs_protocolModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_admissionResults struct {
	AdmissionResult struct {
		ProtocolCommissionMember zfcs_commissionMemberInAppType  `xml:"protocolCommissionMember"`
		Admitted                 bool                            `xml:"admitted"`
		AppRejectedReason        zfcs_appRejectedReasonNotIDType `xml:"appRejectedReason"`
	} `xml:"admissionResult"`
}

type zfcs_applicationAdmittedInfoType struct {
}

type zfcs_applicationCorrespondence struct {
	Compatible   bool    `xml:"compatible"`   /* Флаг соответствия */
	OverallValue float64 `xml:"overallValue"` /* Общая величина преимущества заявки (не используется) */
}

type zfcs_appRejectedReasonType struct {
	NsiRejectReason struct {
		Id     int64  `xml:"id"`
		Reason string `xml:"reason"`
	} `xml:"nsiRejectReason"`
	Explanation zfcs_longTextType `xml:"explanation"` /* Объяснение */
}

type zfcs_appRejectedReasonNotIDType struct {
	NsiRejectReason struct {
		Id     int64  `xml:"id"`
		Reason string `xml:"reason"`
	} `xml:"nsiRejectReason"`
	Explanation zfcs_longTextType `xml:"explanation"` /* Объяснение */
}

type zfcs_bidType struct {
	Price                zfcs_moneyType `xml:"price"`                /* Предложение цены */
	Date                 string         `xml:"date"`                 /* Дата и время подачи ценового предложения */
	IncreaseInitialPrice bool           `xml:"increaseInitialPrice"` /* Признак ценового предложения на повышение начальной (максимальной) цены контракта */
}

type zfcs_commissionMemberInAppType struct {
	MemberNumber uint64 `xml:"memberNumber"` /* Порядковый номер члена комиссии */
}

type zfcs_commissionMemberType struct {
	MemberNumber uint64                  `xml:"memberNumber"` /* Порядковый номер члена комиссии */
	LastName     string                  `xml:"lastName"`     /* Фамилия */
	FirstName    string                  `xml:"firstName"`    /* Имя */
	MiddleName   string                  `xml:"middleName"`   /* Отчество */
	Role         zfcs_commissionRoleType `xml:"role"`         /* Роль члена комиссии */
}

type zfcs_commissionRoleType struct {
	Id        int64  `xml:"id"`        /* Идентификатор */
	Name      string `xml:"name"`      /* Наименование роли */
	RightVote bool   `xml:"rightVote"` /* Имеет право голоса */
}

type zfcs_commissionType struct {
	CommissionName    zfcs_longTextType `xml:"commissionName"` /* Название комиссии */
	CommissionMembers struct {
		CommissionMember          `xml:"commissionMember"`
		SpelledMembersNoVoteCount zfcs_longTextType `xml:"spelledMembersNoVoteCount"`
		SpelledMembersCount       zfcs_longTextType `xml:"spelledMembersCount"`
	} `xml:"commissionMembers"`
	Competent bool              `xml:"competent"` /* Комиссия правомочна осуществлять свои функции в соответствии с Федеральным законом №44-ФЗ */
	AddInfo   zfcs_longTextType `xml:"addInfo"`   /* Дополнительная информация */
}

type zfcs_criterionType struct {
	Code zfcs_criterionCodeEnum `xml:"code"` /* Код критерия:
	CP - Цена контракта.
	MC - Расходы на эксплуатацию и ремонт товаров, использование результатов работ.
	TC - Стоимость жизненного цикла товара или созданного в результате выполнения работы объекта.
	EN - Предложение о сумме соответствующих расходов заказчика, которые заказчик осуществит или понесет по энергосервисному контракту.
	QF - Качественные, функциональные и экологические характеристики объекта закупки.
	QO - Квалификация участников закупки, в том числе наличие у них финансовых ресурсов, на праве собственности или ином законном основании оборудования и других материальных ресурсов, опыта работы, связанного с предметом контракта, и деловой репутации, специалистов и иных работников определенного уровня */
	Value      zfcs_valueType    `xml:"value"`   /* Значимость критерия  */
	AddInfo    zfcs_longTextType `xml:"addInfo"` /* Дополнительная информация о содержании и порядке оценки по критерию */
	Indicators struct {
		Indicator zfcs_indicatorType `xml:"indicator"`
	} `xml:"indicators"`
	Limit            zfcs_indicatorValueType   `xml:"limit"`            /* Предельное значение критерия */
	MeasurementOrder zfcs_measurementOrderEnum `xml:"measurementOrder"` /* Порядок оценки по критерию:
	F - лучшим условием исполнения контракта является наибольшее значение,
	L - лучшим условием исполнения контракта является наименьшее значение,
	O - оценка производится по шкале оценки или другому порядку, указанному в документации */
}

type zfcs_documentRequirementType struct {
	Number    uint64            `xml:"number"`    /* Порядковый номер информации или документа */
	Name      zfcs_longTextType `xml:"name"`      /* Наименование информации или документа */
	Mandatory bool              `xml:"mandatory"` /* Обзательность предоставления */
}

type zfcs_foundationProtocolInfoType struct {
	Name  zfcs_longTextType `xml:"name"`  /* Полное наименование предыдущего протокола */
	Date  string            `xml:"date"`  /* Дата составления предыдущего протокола */
	Place zfcs_longTextType `xml:"place"` /* Место проведения процедуры предыдущего протокола */
}

type zfcs_indicatorType struct {
	Code             int64                     `xml:"code"`             /* Код показателя */
	Name             zfcs_longTextType         `xml:"name"`             /* Наименование показателя */
	Value            zfcs_valueType            `xml:"value"`            /* Значимость показателя */
	Limit            zfcs_indicatorValueType   `xml:"limit"`            /* Предельное значение показателя */
	MeasurementOrder zfcs_measurementOrderEnum `xml:"measurementOrder"` /* Порядок оценки по показателю:
	F - лучшим условием исполнения контракта является наибольшее значение,
	L - лучшим условием исполнения контракта является наименьшее значение,
	O - оценка производится по шкале оценки или другому порядку, указанному в документации */
}

type zfcs_participantType struct {
	ParticipantType zfcs_participantTypeEnum `xml:"participantType"` /* Тип участника:
	P - Физическое лицо РФ;
	PF - Физическое лицо иностранного государства;
	U - Юридическое лицо РФ;
	UF - Юридическое лицо иностранного государства;
	B - Индивидуальный предприниматель РФ;
	BF - Индивидуальный предприниматель иностранного государства. */
	Inn               zfcs_innType               `xml:"inn"`               /* ИНН */
	Kpp               zfcs_kppType               `xml:"kpp"`               /* КПП */
	Ogrn              zfcs_ogrnType              `xml:"ogrn"`              /* ОГРН */
	LegalForm         zfcs_okopfRef              `xml:"legalForm"`         /* Организационно-правовая форма организации в ОКОПФ */
	IdNumber          string                     `xml:"idNumber"`          /* Идентификационный номер для физического или юридического лица иностранного государства */
	IdNumberExtension string                     `xml:"idNumberExtension"` /* Дополнительный идентификационный номер для физического или юридического лица иностранного государства */
	OrganizationName  zfcs_longTextType          `xml:"organizationName"`  /* Наименование организации */
	FirmName          zfcs_longTextType          `xml:"firmName"`          /* Фирменное наименование поставщика */
	Country           zfcs_countryRef            `xml:"country"`           /* Код страны в ОКСМ */
	FactualAddress    string                     `xml:"factualAddress"`    /* Фактический адрес */
	PostAddress       string                     `xml:"postAddress"`       /* Почтовый адрес */
	ContactInfo       zfcs_contactPersonType     `xml:"contactInfo"`       /* Контактная информация */
	ContactEMail      string                     `xml:"contactEMail"`      /* E-mail адрес контактного лица */
	ContactPhone      zfcs_string                `xml:"contactPhone"`      /* Телефон контактного лица */
	ContactFax        zfcs_string                `xml:"contactFax"`        /* Факс контактного лица */
	AdditionalInfo    string                     `xml:"additionalInfo"`    /* Дополнительная информация */
	Status            zfcs_participantStatusType `xml:"status"`            /* Статусы поставщика (исполнителя, подрядчика):
	1 - субъект малого предпринимательства;
	2 - учреждение и предприятие уголовно-исполнительной системы;
	3 - организация инвалидов;
	4 - социально ориентированная некоммерческая организация. */
}

type zfcs_protocolCancelReasonType struct {
}

type zfcs_protocolCancelType struct {
	Id              int64                         `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId      zfcs_externalIdType           `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber  zfcs_purchaseNumberType       `xml:"purchaseNumber"` /* Номер закупки */
	ProtocolNumber  zfcs_documentNumberType       `xml:"protocolNumber"` /* Номер отменяемого протокола */
	DocNumber       zfcs_documentNumberType       `xml:"docNumber"`      /* Номер документа */
	DocDate         string                        `xml:"docDate"`        /* Дата создания документа */
	DocPublishDate  string                        `xml:"docPublishDate"` /* Дата публикации документа */
	Href            zfcs_hrefType                 `xml:"href"`           /* Гиперссылка на опубликованный документ */
	PrintForm       zfcs_printFormType            `xml:"printForm"`      /* Печатная форма документа */
	ExtPrintForm    zfcs_extPrintFormType         `xml:"extPrintForm"`   /* Электронный документ, полученный из внешней системы */
	AddInfo         zfcs_longTextType             `xml:"addInfo"`        /* Дополнительная информация */
	Attachments     zfcs_attachmentListType       `xml:"attachments"`    /* Информация о прикрепленных документах */
	CancelReason    zfcs_protocolCancelReasonType `xml:"cancelReason"`   /* Причина отмены протокола */
	CancelProtocols struct {
		CancelProtocol `xml:"cancelProtocol"`
	} `xml:"cancelProtocols"`
	CancelOrg struct {
		CancelOrg     zfcs_purchaseOrganizationType `xml:"cancelOrg"`
		CancelOrgRole zfcs_responsibleRoleType      `xml:"cancelOrgRole"`
	} `xml:"cancelOrg"`
}

type zfcs_protocolEF1Type struct {
	BasedOn     zfcs_purchaseProtocolEFType
	ProtocolLot struct {
		Applications    `xml:"applications"`     /* Заявки по лоту */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся */
	} `xml:"protocolLot"`
}

type zfcs_protocolEF2Type struct {
	BasedOn     zfcs_purchaseProtocolEFNoCommissionType
	ProtocolLot struct {
		Applications `xml:"applications"` /* Заявки по лоту */
	} `xml:"protocolLot"`
}

type zfcs_protocolEF3Type struct {
	BasedOn     zfcs_purchaseProtocolEFType
	ProtocolLot struct {
		Applications    `xml:"applications"`     /* Заявки по лоту */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся */
	} `xml:"protocolLot"`
}

type zfcs_protocolEFInvalidationType struct {
	BasedOn     zfcs_purchaseProtocolEFNoCommissionType
	ProtocolLot struct {
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся */
	} `xml:"protocolLot"`
}

type zfcs_protocolEFSingleAppType struct {
	BasedOn     zfcs_purchaseProtocolEFType
	ProtocolLot struct {
		Application     `xml:"application"`      /* Заявка по лоту */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся */
	} `xml:"protocolLot"`
	AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся (Устарело) */
}

type zfcs_protocolEFSinglePartType struct {
	BasedOn     zfcs_purchaseProtocolEFType
	ProtocolLot struct {
		Application     `xml:"application"`      /* Заявка по лоту */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся */
	} `xml:"protocolLot"`
	AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание аукциона несостоявшимся (Устарело) */
}

type zfcs_protocolEvasionType struct {
	Id                       int64                   `xml:"id"`                       /* Идентификатор документа ЕИС */
	ExternalId               zfcs_externalIdType     `xml:"externalId"`               /* Внешний идентификатор документа */
	PurchaseNumber           zfcs_purchaseNumberType `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber zfcs_documentNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	FoundationProtocolName   zfcs_longTextType       `xml:"foundationProtocolName"`   /* Наименование предыдущего протокола */
	ParentProtocolNumber     zfcs_documentNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    zfcs_longTextType       `xml:"place"`                    /* Место проведения процедуры */
	ProtocolDate             string                  `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string                  `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string                  `xml:"publishDate"`              /* Дата публикации */
	Href                     zfcs_hrefType           `xml:"href"`                     /* Гиперссылка на опубликованный документ */
	PrintForm                zfcs_printFormType      `xml:"printForm"`                /* Печатная форма протокола */
	ExtPrintForm             zfcs_extPrintFormType   `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	ProtocolPublisher        struct {
		PublisherOrg  zfcs_purchaseOrganizationType `xml:"publisherOrg"`
		PublisherRole zfcs_responsibleRoleType      `xml:"publisherRole"`
	} `xml:"protocolPublisher"`
	PurchaseInfo struct {
		PurchaseResponsible    `xml:"purchaseResponsible"`
		PlacingWay             zfcs_placingWayType     `xml:"placingWay"`
		PublishDate            string                  `xml:"publishDate"`
		PurchaseObjectInfo     zfcs_longTextType       `xml:"purchaseObjectInfo"`
		NotificationFullNumber zfcs_documentNumberType `xml:"notificationFullNumber"`
		NotificationFullName   zfcs_longTextType       `xml:"notificationFullName"`
	} `xml:"purchaseInfo"`
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	Modification zfcs_protocolModificationType `xml:"modification"` /* Основание внесения изменений */
	ProtocolLot  struct {
		LotNumber   uint64           `xml:"lotNumber"`
		LotInfo     zfcs_lotInfoType `xml:"lotInfo"`
		Application `xml:"application"`
		Customer    zfcs_organizationRef `xml:"customer"`
		RefusalFact zfcs_refusalFact     `xml:"refusalFact"`
	} `xml:"protocolLot"`
}

type zfcs_protocolDeviationType struct {
	Id                       int64                   `xml:"id"`                       /* Идентификатор документа ЕИС */
	ExternalId               zfcs_externalIdType     `xml:"externalId"`               /* Внешний идентификатор документа */
	PurchaseNumber           zfcs_purchaseNumberType `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber zfcs_documentNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	FoundationProtocolName   zfcs_longTextType       `xml:"foundationProtocolName"`   /* Наименование предыдущего протокола */
	ParentProtocolNumber     zfcs_documentNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    zfcs_longTextType       `xml:"place"`                    /* Место проведения процедуры */
	ProtocolDate             string                  `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string                  `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string                  `xml:"publishDate"`              /* Дата публикации */
	ExtPrintForm             zfcs_extPrintFormType   `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	Commission               zfcs_commissionType     `xml:"commission"`               /* Информация о комиссии */
	Href                     zfcs_hrefType           `xml:"href"`                     /* Гиперссылка на опубликованный документ */
	PrintForm                zfcs_printFormType      `xml:"printForm"`                /* Печатная форма протокола */
	ProtocolPublisher        struct {
		PublisherOrg  zfcs_purchaseOrganizationType `xml:"publisherOrg"`
		PublisherRole zfcs_responsibleRoleType      `xml:"publisherRole"`
	} `xml:"protocolPublisher"`
	PurchaseInfo struct {
		PurchaseResponsible    `xml:"purchaseResponsible"`
		PlacingWay             zfcs_placingWayType     `xml:"placingWay"`
		PublishDate            string                  `xml:"publishDate"`
		PurchaseObjectInfo     zfcs_longTextType       `xml:"purchaseObjectInfo"`
		NotificationFullNumber zfcs_documentNumberType `xml:"notificationFullNumber"`
		NotificationFullName   zfcs_longTextType       `xml:"notificationFullName"`
	} `xml:"purchaseInfo"`
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	Modification zfcs_protocolModificationType `xml:"modification"` /* Основание внесения изменений */
	ProtocolLot  struct {
		LotNumber   uint64           `xml:"lotNumber"`
		LotInfo     zfcs_lotInfoType `xml:"lotInfo"`
		Application `xml:"application"`
		Customer    zfcs_organizationRef `xml:"customer"`
		RefusalFact `xml:"refusalFact"`
	} `xml:"protocolLot"`
}

type zfcs_protocolModificationReasonType struct {
}

type zfcs_protocolModificationType struct {
	ModificationNumber int                                 `xml:"modificationNumber"` /* Номер изменения */
	Info               zfcs_longTextType                   `xml:"info"`               /* Краткое описание изменения */
	AddInfo            zfcs_longTextType                   `xml:"addInfo"`            /* Дополнительная информация */
	Reason             zfcs_protocolModificationReasonType `xml:"reason"`             /* Основание внесения изменений */
}

type zfcs_protocolOK1Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
}

type zfcs_protocolOK2Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола (для печатной формы) */
}

type zfcs_protocolOKSingleAppType struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола (для печатной формы) */
}

type zfcs_protocolOKD1Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	IsPrequalification bool `xml:"isPrequalification"` /* Предквалифиционный отбор (для печатной формы) */
}

type zfcs_protocolOKD2Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола (для печатной формы) */
}

type zfcs_protocolOKD3Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol `xml:"foundationProtocol"`      /* Реквизиты предыдущего протокола (для печатной формы) */
	OpeningProtocol    zfcs_foundationProtocolInfoType `xml:"openingProtocol"` /* Реквизиты протокола вскрытия первоначальных заявок (для печатной формы) */
}

type zfcs_protocolOKD4Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола */
}

type zfcs_protocolOKD5Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола */
}

type zfcs_protocolOKDSingleAppType struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола */
}

type zfcs_protocolOKOU1Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
}

type zfcs_protocolOKOU2Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола (для печатной формы) */
}

type zfcs_protocolOKOU3Type struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола */
	OpeningProtocol    zfcs_foundationProtocolInfoType `xml:"openingProtocol"`    /* Реквизиты протокола вскрытия заявок (для печатной формы) */
}

type zfcs_protocolOKOUSingleAppType struct {
	BasedOn      zfcs_purchaseProtocolType
	ProtocolLots struct {
		ProtocolLot `xml:"protocolLot"` /* Лот протокола */
	} `xml:"protocolLots"`
	FoundationProtocol zfcs_foundationProtocolInfoType `xml:"foundationProtocol"` /* Реквизиты предыдущего протокола */
	OpeningProtocol    zfcs_foundationProtocolInfoType `xml:"openingProtocol"`    /* Реквизиты протокола вскрытия заявок (для печатной формы) */
}

type zfcs_protocolPOType struct {
	BasedOn     zfcs_purchaseProtocolType
	ProtocolLot struct {
		LotInfo      zfcs_lotInfoType     `xml:"lotInfo"` /* Информация о лоте (для печатной формы) */
		Applications `xml:"applications"` /* Заявки */
	} `xml:"protocolLot"`
}

type zfcs_protocolZKAfterProlongType struct {
	BasedOn     zfcs_purchaseProtocolType
	ProtocolLot struct {
		LotNumber       uint64                   `xml:"lotNumber"` /* Номер лота в извещении */
		Execution       `xml:"execution"`        /* Проведение процедуры */
		LotInfo         zfcs_lotInfoType         `xml:"lotInfo"` /* Информация о лоте (для печатной формы) */
		Applications    `xml:"applications"`     /* Заявки */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание запроса котировок несостоявшимся */
	} `xml:"protocolLot"`
}

type zfcs_protocolZKBIType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	SignDate            string                  `xml:"signDate"`           /* Дата подписания протокола */
	DocNumber           zfcs_documentNumberType `xml:"docNumber"`          /* Номер документа */
	Href                zfcs_hrefType           `xml:"href"`               /* Гиперссылка на опубликованный документ */
	PrintForm           zfcs_printFormType      `xml:"printForm"`          /* Печатная форма документа */
	ExtPrintForm        zfcs_extPrintFormType   `xml:"extPrintForm"`       /* Электронный документ, полученный из внешней системы */
	PurchaseObjectInfo  zfcs_longTextType       `xml:"purchaseObjectInfo"` /* Наименование объекта закупки */
	PurchaseResponsible struct {
		ResponsibleOrg  zfcs_purchaseOrganizationType `xml:"responsibleOrg"`
		ResponsibleRole zfcs_responsibleRoleType      `xml:"responsibleRole"`
		ResponsibleInfo `xml:"responsibleInfo"`
		SpecializedOrg  zfcs_purchaseOrganizationType `xml:"specializedOrg"`
	} `xml:"purchaseResponsible"`
	PlacingWay `xml:"placingWay"` /* Подспособ определения поставщика */
	Lot        struct {
		MaxPrice               zfcs_moneyType                  `xml:"maxPrice"`
		PriceFormula           zfcs_longTextType               `xml:"priceFormula"`
		StandardContractNumber zfcs_standardContractNumberType `xml:"standardContractNumber"`
		Currency               zfcs_currencyRef                `xml:"currency"`
		FinanceSource          zfcs_longTextType               `xml:"financeSource"`
		QuantityUndefined      bool                            `xml:"quantityUndefined"`
		CustomerRequirements   `xml:"customerRequirements"`
		PurchaseObjects        `xml:"purchaseObjects"`
		Preferenses            `xml:"preferenses"`
		Requirements           `xml:"requirements"`
		RestrictInfo           zfcs_longTextType         `xml:"restrictInfo"`
		AddInfo                zfcs_longTextType         `xml:"addInfo"`
		PublicDiscussion       zfcs_publicDiscussionType `xml:"publicDiscussion"`
		NoPublicDiscussion     bool                      `xml:"noPublicDiscussion"`
	} `xml:"lot"`
	Okpd2okved2  bool                              `xml:"okpd2okved2"`  /* Классификация по ОКПД2/ОКВЭД2. Элемент не используется в импорте */
	ProtocolLot  zfcs_infoProtocolZKBIType         `xml:"protocolLot"`  /* Информация о протоколе */
	Attachments  zfcs_attachmentListType           `xml:"attachments"`  /* Документация о закупке */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_protocolZKBIAfterProlongType struct {
	Id             int64                   `xml:"id"`             /* Идентификатор документа ЕИС */
	ExternalId     zfcs_externalIdType     `xml:"externalId"`     /* Внешний идентификатор документа */
	PurchaseNumber zfcs_purchaseNumberType `xml:"purchaseNumber"` /* Номер закупки */
	DocPublishDate string                  `xml:"docPublishDate"` /* Дата публикации документа
	Планируемая или фактическая */
	SignDate     string                            `xml:"signDate"`     /* Дата подписания протокола */
	DocNumber    zfcs_documentNumberType           `xml:"docNumber"`    /* Номер документа */
	Href         zfcs_hrefType                     `xml:"href"`         /* Гиперссылка на опубликованный документ */
	PrintForm    zfcs_printFormType                `xml:"printForm"`    /* Печатная форма документа */
	ExtPrintForm zfcs_extPrintFormType             `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
	ProtocolLot  zfcs_infoProtocolZKBIType         `xml:"protocolLot"`  /* Информация о протоколе */
	Modification zfcs_notificationModificationType `xml:"modification"` /* Основание внесения изменений */
}

type zfcs_infoProtocolZKBIType struct {
	Execution struct {
		Place       zfcs_longTextType `xml:"place"`
		OpeningDate string            `xml:"openingDate"`
		ScoringDate string            `xml:"scoringDate"`
		SignDate    string            `xml:"signDate"`
	} `xml:"execution"`
	Commission   `xml:"commission"` /* Информация о комиссии */
	PurchaseInfo struct {
		PurchaseResponsible    `xml:"purchaseResponsible"`
		PlacingWay             zfcs_placingWayType     `xml:"placingWay"`
		PublishDate            string                  `xml:"publishDate"`
		PurchaseObjectInfo     zfcs_longTextType       `xml:"purchaseObjectInfo"`
		NotificationFullNumber zfcs_documentNumberType `xml:"notificationFullNumber"`
		NotificationFullName   zfcs_longTextType       `xml:"notificationFullName"`
	} `xml:"purchaseInfo"`
	LotInfo      zfcs_lotInfoType `xml:"lotInfo"` /* Информация о лоте (для печатной формы) */
	Applications struct {
		Application `xml:"application"`
	} `xml:"applications"`
	AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание запроса котировок несостоявшимся */
}

type zfcs_protocolZKType struct {
	BasedOn     zfcs_purchaseProtocolType
	ProtocolLot struct {
		LotNumber       uint64                   `xml:"lotNumber"` /* Номер лота в извещении */
		Execution       `xml:"execution"`        /* Проведение процедуры */
		LotInfo         zfcs_lotInfoType         `xml:"lotInfo"` /* Информация о лоте (для печатной формы) */
		Applications    `xml:"applications"`     /* Заявки */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание запроса котировок несостоявшимся */
	} `xml:"protocolLot"`
}

type zfcs_protocolZPExtractType struct {
	Id                       int64                   `xml:"id"`                       /* Идентификатор документа ЕИС */
	ExternalId               zfcs_externalIdType     `xml:"externalId"`               /* Внешний идентификатор документа */
	PurchaseNumber           zfcs_purchaseNumberType `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType `xml:"protocolNumber"`           /* Номер протокола */
	FoundationProtocolNumber zfcs_documentNumberType `xml:"foundationProtocolNumber"` /* Номер предыдущего протокола */
	ParentProtocolNumber     zfcs_documentNumberType `xml:"parentProtocolNumber"`     /* Номер родительского протокола - в случае внесения изменений */
	Place                    zfcs_longTextType       `xml:"place"`                    /* Место проведения процедуры */
	ProtocolDate             string                  `xml:"protocolDate"`             /* Дата составления протокола */
	SignDate                 string                  `xml:"signDate"`                 /* Дата подписания протокола */
	PublishDate              string                  `xml:"publishDate"`              /* Дата публикации */
	CommissionName           zfcs_longTextType       `xml:"commissionName"`           /* Название комиссии */
	PrintForm                zfcs_printFormType      `xml:"printForm"`                /* Печатная форма протокола */
	ExtPrintForm             zfcs_extPrintFormType   `xml:"extPrintForm"`             /* Электронный документ, полученный из внешней системы */
	ProtocolPublisher        struct {
		PublisherOrg  zfcs_purchaseOrganizationType `xml:"publisherOrg"`
		PublisherRole zfcs_responsibleRoleType      `xml:"publisherRole"`
	} `xml:"protocolPublisher"`
	PurchaseInfo struct {
		PurchaseResponsible  `xml:"purchaseResponsible"`
		PlacingWay           zfcs_placingWayType `xml:"placingWay"`
		PublishDate          string              `xml:"publishDate"`
		PurchaseObjectInfo   zfcs_longTextType   `xml:"purchaseObjectInfo"`
		NotificationFullName zfcs_longTextType   `xml:"notificationFullName"`
	} `xml:"purchaseInfo"`
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	Modification zfcs_protocolModificationType `xml:"modification"` /* Основание внесения изменений */
	Href         zfcs_hrefType                 `xml:"href"`         /* Гиперссылка на опубликованный документ */
	ProtocolLot  struct {
		LotInfo          zfcs_lotInfoType `xml:"lotInfo"`
		Applications     `xml:"applications"`
		BestPrice        zfcs_moneyType           `xml:"bestPrice"`
		SpelledBestPrice zfcs_longTextType        `xml:"spelledBestPrice"`
		AbandonedReason  zfcs_abandonedReasonType `xml:"abandonedReason"`
	} `xml:"protocolLot"`
}

type zfcs_protocolZPFinalType struct {
	BasedOn     zfcs_purchaseProtocolType
	ProtocolLot struct {
		LotInfo         zfcs_lotInfoType         `xml:"lotInfo"` /* Информация о лоте (для печатной формы) */
		Applications    `xml:"applications"`     /* Заявки */
		AbandonedReason zfcs_abandonedReasonType `xml:"abandonedReason"` /* Признание запроса предложений несостоявшимся (для печатной формы) */
	} `xml:"protocolLot"`
	FoundationProtocolName zfcs_longTextType `xml:"foundationProtocolName"` /* Наименование предыдущего протокола (для печатной формы) */
}

type zfcs_protocolZPType struct {
	BasedOn     zfcs_purchaseProtocolType
	ProtocolLot struct {
		LotInfo              zfcs_lotInfoType             `xml:"lotInfo"`      /* Информация о лоте (для печатной формы) */
		NoLastOffers         bool                         `xml:"noLastOffers"` /* Все участники отказались направить окончательные предложения. Все заявки, соответствующие требованиям признаны окончательными предложениями (согласно ч.14 ст. 83 Федерального закона 44-ФЗ) */
		Applications         `xml:"applications"`         /* Допущеные заявки */
		RejectedApplications `xml:"rejectedApplications"` /* Отклоненные заявки (для печатной формы) */
		AbandonedReason      zfcs_abandonedReasonType     `xml:"abandonedReason"` /* Признание запроса предложений несостоявшимся */
	} `xml:"protocolLot"`
}

type zfcs_auditActionSubjectsRef struct {
	Id   int64                `xml:"id"`   /* Идентификатор предмета */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование предмета */
}

type zfcs_countryRef struct {
	CountryCode     string `xml:"countryCode"`     /* Цифровой код страны */
	CountryFullName string `xml:"countryFullName"` /* Полное наименование страны */
}

type zfcs_currencyRef struct {
	Code string `xml:"code"` /* Код валюты */
	Name string `xml:"name"` /* Наименование валюты */
}

type zfcs_currencyFullRef struct {
	Code        string `xml:"code"`        /* Код валюты */
	DigitalCode string `xml:"digitalCode"` /* Цифровой код валюты */
	Name        string `xml:"name"`        /* Наименование валюты */
}

type zfcs_OKEIRef struct {
	Code string `xml:"code"` /* Код */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_okopfRef struct {
	Code         string               `xml:"code"`         /* Код */
	SingularName zfcs_longTextMinType `xml:"singularName"` /* Наименование в единственном числе */
}

type zfcs_OKPDRef struct {
	Code string `xml:"code"` /* Код товара, работы или услуги */
	Name string `xml:"name"` /* Наименование товара, работы или услуги */
}

type zfcs_OKTMORef struct {
	Code string `xml:"code"` /* Код по ОКТМО */
	Name string `xml:"name"` /* Полное наименование */
}

type zfcs_OKVEDRef struct {
	Code string `xml:"code"` /* Код */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_OKATORef struct {
	Code string            `xml:"code"` /* Код по ОKATO */
	Name zfcs_longTextType `xml:"name"` /* Наименование */
}

type zfcs_OKPORef struct {
	Code zfcs_okpoType        `xml:"code"` /* Код */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование */
}

type zfcs_OKFSRef struct {
	Code string               `xml:"code"` /* Код формы собственности */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование */
}

type zfcs_OKSMRef struct {
	CountryCode     string `xml:"countryCode"`     /* Цифровой код страны */
	CountryFullName string `xml:"countryFullName"` /* Полное наименование страны */
}

type zfcs_subjectRFRef struct {
	Code string `xml:"code"` /* Код субъекта */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_PPORef struct {
	Code zfcs_ppoCodeType `xml:"code"` /* Код публично-правового образования */
	Name string           `xml:"name"` /* Наименование публично-правового образования */
}

type zfcs_contractTerminationReasonType struct {
	Code zfcs_nsiContractTerminationReasonCodeType `xml:"code"` /* Код */
	Name zfcs_nsiContractTerminationReasonNameType `xml:"name"` /* Наименование основания */
}

type zfcs_organizationLink struct {
	Id          int64                                `xml:"id"`          /* Идентификатор связи */
	ActiveUntil string                               `xml:"activeUntil"` /* Срок окончания действия права (отсутствует если связь не имеет срока действия) */
	BlockStatus zfcs_organizationLinkBlockStatusType `xml:"blockStatus"` /* Статус права

	A - Активно
	B - Заблокировано
	BO - Заблокировано создание новых извещений */
	DependentOrganization zfcs_organizationRef                 `xml:"dependentOrganization"` /* Подчиненная организация (например, Заказчик в случае связи СО - Заказчик/УО) или Главная организация, для которой создается связь (например, Специализированная организация в случае связи СО - Заказчик/УО) */
	LinkUsers             zfcs_linkUser                        `xml:"linkUsers"`             /* Список связей "право - пользователь". Список пользователей, наделенными данными полномочиями */
	OrdersVisibilityType  zfcs_orgLinkOrdersVisibilityTypeEnum `xml:"ordersVisibilityType"`  /* Область доступа к осуществляемым закупкам:

	A - Активно;
	B - Заблокировано.				 */
	LastModifyDate string `xml:"lastModifyDate"` /* Дата и время последнего изменения */
}

type zfcs_organizationRef struct {
	RegNum          zfcs_string              `xml:"regNum"`          /* Код по СПЗ */
	ConsRegistryNum zfcs_consRegistryNumType `xml:"consRegistryNum"` /* Код по Сводному Реестру */
	FullName        zfcs_longTextMinType     `xml:"fullName"`        /* Полное наименование */
}

type zfcs_publicDiscussionFoundationRef struct {
	Code string               `xml:"code"` /* Код основания результата общественного обсуждения */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование основания */
}

type zfcs_publicDiscussionDecisionRef struct {
	Code string               `xml:"code"` /* Код решения общественного обсуждения */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование решения */
}

type zfcs_publicDiscussionFacetRef struct {
	Code      string               `xml:"code"`      /* Код аспекта обсуждений */
	FacetName zfcs_longTextMinType `xml:"facetName"` /* Наименование аспекта */
}

type zfcs_publicDiscussionQuestionRef struct {
	Code string               `xml:"code"` /* Код вопроса */
	Name zfcs_longTextMinType `xml:"name"` /* Текст вопроса */
}

type zfcs_etpPrivilege struct {
	Etp       zfcs_ETPType       `xml:"etp"`       /* ЭП */
	EtpAction zfcs_etpActionType `xml:"etpAction"` /* Действия пользователя на ЭП:

	PRP - Размещение результатов рассмотрения первых частей заявок;
	PR1 - Размещение протокола подведения итогов;
	PR2- Размещение протокола отказа от заключения контракта;
	PR3 - Подписание государственного (муниципального) контракта;
	PR4 - Направление проекта контракта участнику размещения заказа. */
	Organization zfcs_organizationRef     `xml:"organization"` /* Для какой организации. Если не указан, то для своей */
	Status       zfcs_etpActionStatusType `xml:"status"`       /* Статус привилегии

	A - Активно;
	B - Заблокировано;
	RA - Заявка на добавление;
	RB - Заявка на блкировку. */
}

type zfcs_linkUser struct {
	User        string                       `xml:"user"`        /* Логин пользователя */
	BlockStatus zfcs_userLinkBlockStatusType `xml:"blockStatus"` /* Статус права для пользователя:

	A - Активно;
	B - Заблокировано;
	BO - Заблокировано создание новых извещений. */
}

type zfcs_masterDataType struct {
}

type zfcs_nsiAbandonedReasonType struct {
	Id         int64                        `xml:"id"`         /* Идентификатор основания признания процедуры несостоявшейся */
	Code       string                       `xml:"code"`       /* Код основания признания процедуры несостоявшейся */
	Name       string                       `xml:"name"`       /* Наименование основания признания процедуры несостоявшейся */
	ObjectName string                       `xml:"objectName"` /* Наименование интеграционного объекта, к которому применимо данное основание */
	Type       zfcs_abandonedReasonTypeEnum `xml:"type"`       /* Тип основания:
	OR - По окончании срока подачи заявок подана только одна заявка. Такая заявка признана соответствующей требованиям 44-ФЗ и требованиям, указанным в извещении;
	NR - По окончании срока подачи заявок не подано ни одной заявки;
	OV - По результатам рассмотрения заявок только одна заявка признана соответствующей требованиям ФЗ и требованиям, указанным в извещении;
	NV - Все поданные заявки отклонены;
	OV2 - По результатам рассмотрения вторых частей заявок только одна заявка признана соответствующей требованиям 44-ФЗ и требованиям, указанным в извещении или ни одной заявки не признано соответствующим данным требованиям. */
	DocType    zfcs_docType        `xml:"docType"`    /* Тип документа */
	PlacingWay zfcs_placingWayType `xml:"placingWay"` /* Подспособ определения поставщика */
	Actual     bool                `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiAuditActionSubjectsType struct {
	Id     int64                `xml:"id"`     /* Идентификатор БД */
	Name   zfcs_longTextMinType `xml:"name"`   /* Наименование предмета */
	Actual bool                 `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiBankGuaranteeRefusalReasonType struct {
	Id     int64  `xml:"id"`     /* Идентификатор БД */
	Name   string `xml:"name"`   /* Наименование причины */
	Actual bool   `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiBudgetType struct {
	Code   string `xml:"code"`   /* Код бюджета */
	Name   string `xml:"name"`   /* Наименование */
	Actual bool   `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiCalendarDaysType struct {
	StartDate string `xml:"startDate"` /* Дата начала периода выгрузки календаря */
	EndDate   string `xml:"endDate"`   /* Дата окончания периода выгрузки календаря */
	Days      struct {
		Day `xml:"day"`
	} `xml:"days"`
}

type zfcs_nsiCommissionRoleType struct {
	Id        int64  `xml:"id"`        /* Идентификатор */
	Name      string `xml:"name"`      /* Наименование роли */
	Order     int    `xml:"order"`     /* Порядковый номер */
	RightVote bool   `xml:"rightVote"` /* Имеет право голоса */
	Actual    bool   `xml:"actual"`    /* Признак актуальности */
}

type zfcs_nsiCommissionType struct {
	RegNumber         int    `xml:"regNumber"`      /* Номер в реестре комиссий */
	CommissionName    string `xml:"commissionName"` /* Название комиссии */
	CommissionMembers struct {
		CommissionMember `xml:"commissionMember"`
	} `xml:"commissionMembers"`
	Owner     zfcs_organizationRef `xml:"owner"`     /* Организация, размещающая заказ */
	Competent bool                 `xml:"competent"` /* Комиссия правомочна осуществлять свои функции в соответствии с Федеральным законом №44-ФЗ */
	AddInfo   zfcs_longTextType    `xml:"addInfo"`   /* Дополнительная информация */
	Actual    bool                 `xml:"actual"`    /* Признак актуальности */
}

type zfcs_nsiContractPriceChangeReasonType struct {
	Id            int64                  `xml:"id"`            /* Идентификатор БД */
	Name          string                 `xml:"name"`          /* Наименование обоснования */
	SubsystemType zfcs_subsystemTypeEnum `xml:"subsystemType"` /* К какому закону относится обоснование:

	FZ94 - 94ФЗ:
	FZ44 - 44ФЗ. */
	Actual bool `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiContractRefusalReasonType struct {
	Id     int64  `xml:"id"`     /*  */
	Name   string `xml:"name"`   /* Наименование основания */
	Actual bool   `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiContractSingleCustomerReasonType struct {
	Id            int64                                `xml:"id"`            /* Идентификатор БД (Устарело. с 01.01.2015 не применяется) */
	Code          zfcs_nsiSingleCustomerReasonCodeType `xml:"code"`          /* Код */
	Name          zfcs_nsiSingleCustomerReasonNameType `xml:"name"`          /* Наименование основания */
	PointLaw      string                               `xml:"pointLaw"`      /* Пункт и статья закона */
	SubsystemType zfcs_subsystemTypeEnum               `xml:"subsystemType"` /* К какому закону относится основание:

	FZ94 - 94ФЗ:
	FZ44 - 44ФЗ. */
	Actual    bool `xml:"actual"` /* Признак актуальности */
	Documents struct {
		Document `xml:"document"`
	} `xml:"documents"`
}

type zfcs_nsiContractTerminationReasonType struct {
	Id            int64                                     `xml:"id"`            /* Идентификатор БД. (Устарело. с 01.01.2015 не применяется) */
	Code          zfcs_nsiContractTerminationReasonCodeType `xml:"code"`          /* Код */
	Name          zfcs_nsiContractTerminationReasonNameType `xml:"name"`          /* Наименование основания */
	SubsystemType zfcs_subsystemTypeEnum                    `xml:"subsystemType"` /* К какому закону относится основание:

	FZ94 - 94ФЗ:
	FZ44 - 44ФЗ. */
	Actual    bool `xml:"actual"` /* Признак актуальности */
	Documents struct {
		Document `xml:"document"`
	} `xml:"documents"`
	ReparationsDocuments struct {
		Document `xml:"document"`
	} `xml:"reparationsDocuments"`
}

type zfcs_nsiContractModificationReasonType struct {
	Code      zfcs_nsiContractModificationReasonCodeType `xml:"code"`   /* Код */
	Name      zfcs_nsiContractModificationReasonNameType `xml:"name"`   /* Наименование причины */
	Actual    bool                                       `xml:"actual"` /* Признак актуальности */
	Documents struct {
		Document `xml:"document"`
	} `xml:"documents"`
}

type zfcs_nsiContractExecutionDocType struct {
	Code   zfcs_nsiContractExecutionDocCodeType `xml:"code"`   /* Код */
	Name   zfcs_nsiContractExecutionDocNameType `xml:"name"`   /* Наименование типа */
	Actual bool                                 `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiContractReparationDocType struct {
	Code   zfcs_nsiContractPenaltyReasonCodeType `xml:"code"`   /* Код */
	Name   zfcs_nsiContractReparationDocNameType `xml:"name"`   /* Наименование типа */
	Actual bool                                  `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiContractPenaltyReasonType struct {
	Code        zfcs_nsiContractPenaltyReasonCodeType `xml:"code"`        /* Код */
	Name        zfcs_nsiContractPenaltyReasonNameType `xml:"name"`        /* Наименование типа */
	PenaltyType zfcs_contractPenaltyType              `xml:"penaltyType"` /* Тип взыскания:
	F - Штраф;
	I - Пени. */
	Actual bool `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiContractOKOPFExtraBudgetType struct {
	Extrabudget  zfcs_extraBudgetFundsContract2015 `xml:"extrabudget"` /* Вид внебюджетных средств */
	LegalFormOld struct {
		Code         string            `xml:"code"`
		SingularName zfcs_longTextType `xml:"singularName"`
	} `xml:"legalFormOld"`
	LegalFormNew okopfType `xml:"legalFormNew"` /* Код и наименование ОКОПФ (ОК 028-2012) */
	Actual       bool      `xml:"actual"`       /* Признак актуальности */
}

type zfcs_nsiContractCurrencyCBRFType struct {
	Currency zfcs_currencyFullRef `xml:"currency"` /* Валюта */
	Actual   bool                 `xml:"actual"`   /* Признак актуальности */
}

type zfcs_nsiCurrencyType struct {
	Code        string `xml:"code"`        /* Код валюты */
	DigitalCode string `xml:"digitalCode"` /* Цифровой код валюты */
	Name        string `xml:"name"`        /* Наименование валюты */
	Actual      bool   `xml:"actual"`      /* Признак актуальности */
}

type zfcs_nsiEvalCriterionType struct {
	Id              int64  `xml:"id"`   /* Идентификатор */
	Name            string `xml:"name"` /* Название критерия */
	CriterionGroups struct {
		CriterionGroup `xml:"criterionGroup"`
	} `xml:"criterionGroups"`
	Code           string `xml:"code"`           /* Символьный код */
	CriterionCode  string `xml:"criterionCode"`  /* Символьный код критерия */
	Description    string `xml:"description"`    /* Описание */
	NumericalCode  int64  `xml:"numericalCode"`  /* Цифровой код */
	Order          int    `xml:"order"`          /* Порядковый номер */
	Actual         bool   `xml:"actual"`         /* Признак актуальности */
	NeedExpertEval bool   `xml:"needExpertEval"` /* Нуждается в экспертной оценке */
}

type zfcs_nsiKBKBudgetType struct {
	Kbk    string `xml:"kbk"`    /* Код КБК */
	Budget string `xml:"budget"` /* Код бюджета */
	Actual bool   `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiKOSGUType struct {
	Code       string `xml:"code"`       /* Код */
	Name       string `xml:"name"`       /* Наименование */
	ParentCode string `xml:"parentCode"` /* Код узла предка в иерархии */
	Actual     bool   `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiOffBudgetType struct {
	Code          string                 `xml:"code"`          /* Код бюджета */
	Name          string                 `xml:"name"`          /* Наименование */
	SubsystemType zfcs_subsystemTypeEnum `xml:"subsystemType"` /* К какому закону относится вид:

	FZ94 - 94ФЗ:
	FZ44 - 44ФЗ. */
	Actual bool `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiOKEIType struct {
	Code     string `xml:"code"`     /* Код */
	FullName string `xml:"fullName"` /* Полное наименование */
	Section  struct {
		Code string `xml:"code"`
		Name string `xml:"name"`
	} `xml:"section"`
	Group struct {
		Id   int64  `xml:"id"`
		Name string `xml:"name"`
	} `xml:"group"`
	LocalName           string `xml:"localName"`           /* Национальное кодовое буквенное обозначение */
	InternationalName   string `xml:"internationalName"`   /* Международное кодовое буквенное обозначение */
	LocalSymbol         string `xml:"localSymbol"`         /* Национальное условное обозначение */
	InternationalSymbol string `xml:"internationalSymbol"` /* Международное условное обозначение */
	Actual              bool   `xml:"actual"`              /* Признак актуальности */
}

type zfcs_nsiOKOPFType struct {
	Code         string `xml:"code"`         /* Код организационно-правовой формы */
	ParentCode   string `xml:"parentCode"`   /* Код узла предка в иерархии */
	FullName     string `xml:"fullName"`     /* Полное наименование */
	SingularName string `xml:"singularName"` /* Наименование в единственном числе */
	Actual       bool   `xml:"actual"`       /* Признак актуальности */
}

type zfcs_nsiOKPDType struct {
	Id         int64  `xml:"id"`         /* Идентификатор в БД */
	ParentId   int64  `xml:"parentId"`   /* Идентификатор родительской записи в БД */
	Code       string `xml:"code"`       /* Код товара, работы или услуги */
	ParentCode string `xml:"parentCode"` /* Код узла предка в иерархии */
	Name       string `xml:"name"`       /* Наименование товара, работы или услуги */
	Comment    string `xml:"comment"`    /* Комментарий */
	Actual     bool   `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiOKPD2Type struct {
	Id         int64             `xml:"id"`         /* Идентификатор в БД */
	ParentId   int64             `xml:"parentId"`   /* Идентификатор родительской записи в БД */
	Code       string            `xml:"code"`       /* Код товара, работы или услуги */
	ParentCode string            `xml:"parentCode"` /* Код узла предка в иерархии */
	Name       zfcs_longTextType `xml:"name"`       /* Наименование товара, работы или услуги */
	Comment    zfcs_longTextType `xml:"comment"`    /* Комментарий */
	Actual     bool              `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiOKSMType struct {
	CountryCode     string `xml:"countryCode"`     /* Цифровой код страны */
	CountryFullName string `xml:"countryFullName"` /* Полное наименование страны */
	Actual          bool   `xml:"actual"`          /* Признак актуальности */
}

type zfcs_nsiOKVEDType struct {
	Id         int64  `xml:"id"`         /* Идентификатор объекта ЕИС */
	Code       string `xml:"code"`       /* Код */
	Section    string `xml:"section"`    /* Код раздела */
	Subsection string `xml:"subsection"` /* Код подраздела */
	ParentId   int64  `xml:"parentId"`   /* Идентификатор узла предка в иерархии */
	Name       string `xml:"name"`       /* Наименование */
	Actual     bool   `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiOKVED2Type struct {
	Id         int64             `xml:"id"`         /* Идентификатор объекта ЕИС */
	Code       string            `xml:"code"`       /* Код */
	Section    string            `xml:"section"`    /* Код раздела */
	ParentCode string            `xml:"parentCode"` /* Код узла предка в иерархии */
	Name       string            `xml:"name"`       /* Наименование */
	Comment    zfcs_longTextType `xml:"comment"`    /* Комментарий */
	Actual     bool              `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiDeviationFactFoundationType struct {
	Code   string `xml:"code"`   /* Код причины признания участника уклонившимся */
	Name   string `xml:"name"`   /* Наименование причины признания участника уклонившимся */
	Actual bool   `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiOKTMOType struct {
	Code       string `xml:"code"`       /* Код */
	ParentCode string `xml:"parentCode"` /* Код узла предка в иерархии */
	Section    string `xml:"section"`    /* Раздел */
	FullName   string `xml:"fullName"`   /* Полное наименование */
	Actual     bool   `xml:"actual"`     /* Признак актуальности */
}

type zfcs_nsiOrganizationRightsType struct {
	RegNumber         zfcs_string `xml:"regNumber"` /* Код по Сводному перечню Заказчиков */
	OrganizationLinks struct {
		OrganizationLink zfcs_organizationLink `xml:"organizationLink"`
	} `xml:"organizationLinks"`
}

type zfcs_nsiOrganizationType struct {
	RegNumber       zfcs_string              `xml:"regNumber"`       /* Реестровый номер в СПЗ */
	ConsRegistryNum zfcs_consRegistryNumType `xml:"consRegistryNum"` /* Код по Сводному Реестру */
	ShortName       string                   `xml:"shortName"`       /* Сокращенное наименование */
	FullName        string                   `xml:"fullName"`        /* Полное наименование */
	GRBSCode        string                   `xml:"GRBSCode"`        /* Код организации по перечню главных распорядителей бюджетных средств */
	BIK             string                   `xml:"BIK"`             /* Банковский идентификационный код в соответствии с классификатором банковских идентификационных кодов, ведущимся Банком России. Указывается в случае если у организации есть полномочие "Банк" */
	NomBank         string                   `xml:"nomBank"`         /* Регистрационный номер банка по Перечню банков, присваиваемый Банком России при регистрации кредитной организации в Книге государственной регистрации кредитных организаций. Указывается в случае если у организации есть полномочие "Банк" */
	FactualAddress  struct {
		OKATO          string          `xml:"OKATO"`
		AddressLine    string          `xml:"addressLine"`
		Area           zfcs_kladrType  `xml:"area"`
		Building       string          `xml:"building"`
		Country        zfcs_countryRef `xml:"country"`
		FilledManually bool            `xml:"filledManually"`
		Office         string          `xml:"office"`
		Region         zfcs_kladrType  `xml:"region"`
		Settlement     zfcs_kladrType  `xml:"settlement"`
		City           zfcs_kladrType  `xml:"city"`
		ShortStreet    string          `xml:"shortStreet"`
		Street         zfcs_kladrType  `xml:"street"`
		Zip            string          `xml:"zip"`
	} `xml:"factualAddress"`
	PostalAddress zfcs_longTextType      `xml:"postalAddress"` /* Почтовый адрес */
	Email         string                 `xml:"email"`         /* Адрес электронной почты для получения системных уведомлений */
	Phone         zfcs_string            `xml:"phone"`         /* Телефон */
	Fax           zfcs_string            `xml:"fax"`           /* Факс */
	ContactPerson zfcs_contactPersonType `xml:"contactPerson"` /* Контактное лицо */
	Accounts      struct {
		Account zfcs_accountType `xml:"account"`
	} `xml:"accounts"`
	Budgets struct {
		Budget `xml:"budget"`
	} `xml:"budgets"`
	HeadAgency       zfcs_organizationRef `xml:"headAgency"`       /* Административная принадлежность */
	OrderingAgency   zfcs_organizationRef `xml:"orderingAgency"`   /* Вышестоящая организация в части подтверждения полномочия в сфере размещения заказа */
	INN              zfcs_innType         `xml:"INN"`              /* Код ИНН */
	KPP              zfcs_kppType         `xml:"KPP"`              /* Код КПП */
	RegistrationDate string               `xml:"registrationDate"` /* Дата постановки на учет в налоговом органе */
	UBPCode          string               `xml:"UBPCode"`          /* Код организации по Сводному реестру участников бюджетного процесса (УБП). Заполняется для УБП федерального бюджета  */
	IKUInfo          struct {
		IKU        zfcs_ikuType `xml:"IKU"`
		DateStIKU  string       `xml:"dateStIKU"`
		DateEndIKU string       `xml:"dateEndIKU"`
	} `xml:"IKUInfo"`
	OGRN  string `xml:"OGRN"` /* Код ОГРН */
	OKOPF struct {
		Code     string `xml:"code"`
		FullName string `xml:"fullName"`
	} `xml:"OKOPF"`
	OKPO  string            `xml:"OKPO"`  /* Код ОКПО */
	OKVED zfcs_longTextType `xml:"OKVED"` /* Код ОКВЭД */
	OKOGU struct {
		Code string `xml:"code"`
		Name string `xml:"name"`
	} `xml:"OKOGU"`
	OrganizationRole zfcs_organizationRoleType `xml:"organizationRole"` /* Полномочия организации:
	CU - Заказчик;
	RA - Уполномоченный орган;
	AI - Уполномоченное учреждение;
	SO - Специализированная организация;
	CO - Контрольный орган;
	SP - Служба Оператора сайта;
	FO - Финансовый орган;
	EO - Оператор электронной площадки;
	AA - Орган аудита;
	CA - Орган по регулированию контрактной системы в сфере закупок;
	NA - Орган, устанавливающий правила нормирования;
	DA - Орган, утверждающий требования к отдельным видам товаров, работ, услуг;
	BA - Банк;
	TA - Орган, разрабатывающий и утверждающий типовые контракты и типовые условия контрактов;
	OA - Организация, осуществляющая полномочия заказчика на осуществление закупок на основании договора (соглашения) в соответствии с частью 6 статьи 15 Закона № 44-ФЗ;
	CIA - Орган контроля соответствия информации об объемах финансового обеспечения и идентификационных кодах закупок;
	ICB - Орган внутреннего контроля;
	NP - Орган, уполномоченный на ведение реестра недобросовестных поставщиков;
	GR - Главный распорядитель бюджетных средств;
	OV - Орган государственной (исполнительной) власти. */
	OrganizationType struct {
		Code string `xml:"code"`
		Name string `xml:"name"`
	} `xml:"organizationType"`
	OKTMO             zfcs_OKTMORef              `xml:"OKTMO"`             /* ОКТМО организации */
	OKTMOPPO          zfcs_OKTMORef              `xml:"OKTMOPPO"`          /* ОКТМО публично-правового образования (ППО) */
	SubordinationType zfcs_subordinationTypeEnum `xml:"subordinationType"` /* Уровень организации:

	1 – Федеральный уровень;
	2 – Уровень субъекта РФ;
	3 – Муниципальный уровень; */
	Url      string `xml:"url"`      /* Адрес сайта в сети Интернет */
	TimeZone int    `xml:"timeZone"` /* Часовая зона заказчика в Российском исчислении часовых зон
	Указывается смещение от Московского времени (MSK) */
	TimeZoneUtcOffset string `xml:"timeZoneUtcOffset"` /* Смещение (в часах и минутах) часовой зоны заказчика относительно UTC. */
	TimeZoneOlson     string `xml:"timeZoneOlson"`     /* Идентификатор часовой зоны заказчика в базе часовых поясов Olson. */
	Actual            bool   `xml:"actual"`            /* Признак актуальности */
	Register          bool   `xml:"register"`          /* Признак регистрации на ЕИС */
}

type zfcs_nsiOrganizationTypesType struct {
	Code        string `xml:"code"`        /* Код типа организации */
	Name        string `xml:"name"`        /* Наименование типа организации */
	Description string `xml:"description"` /* Описание */
}

type zfcs_nsiPlacingWayType struct {
	PlacingWayId  int64                      `xml:"placingWayId"`  /* Идентификатор способа размещения заказа (определения поставщика) */
	Code          zfcs_nsiPlacingWayCodeType `xml:"code"`          /* Подспособ размещения заказа (определения поставщика) */
	Name          string                     `xml:"name"`          /* Наименование подспособа размещения заказа (определения поставщика) */
	Type          string                     `xml:"type"`          /* Способ размещения заказа (определения поставщика) */
	SubsystemType zfcs_placingWayTypeEnum    `xml:"subsystemType"` /* Тип способa размещения заказа (определения поставщика):

	FZ94 - закон 94ФЗ;
	FZ44 - закон 44ФЗ. */
	Actual    bool `xml:"actual"` /* Признак актуальности */
	Documents struct {
		Document `xml:"document"`
	} `xml:"documents"`
}

type zfcs_nsiPlanPositionChangeReasonType struct {
	Id          int64  `xml:"id"`          /* Идентификатор */
	Name        string `xml:"name"`        /* Сокращенное описание обоснования внесения изменений */
	Description string `xml:"description"` /* Подробное описание обоснования изменений */
	Actual      bool   `xml:"actual"`      /* Признак актуальности */
}

type zfcs_nsiPublicDiscussionDecisionsType struct {
	Id          int64                `xml:"id"`   /* Идентификатор решения общественного обсуждения */
	Code        string               `xml:"code"` /* Код решения общественного обсуждения */
	Name        zfcs_longTextMinType `xml:"name"` /* Наименование решения */
	Foundations struct {
		Foundation `xml:"foundation"`
	} `xml:"foundations"`
	Type zfcs_publicDiscussionTypeEnum `xml:"type"` /* Тип обсуждений:
	LP - общественное обсуждение крупных закупок,
	OT - иное общественное обсуждение */
	Phase zfcs_publicDiscussionLargePurchaseStagesEnum `xml:"phase"` /* Этап общественного обсуждения.
	Возможные значения:
	S1 - Этап 1
	S2 - Этап 2 */
	Actual bool `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiPublicDiscussionQuestionnarieType struct {
	Id        int64                `xml:"id"`        /* Идентификатор аспекта обсуждений */
	Code      string               `xml:"code"`      /* Код аспекта обсуждений */
	FacetName zfcs_longTextMinType `xml:"facetName"` /* Наименование аспекта */
	Questions struct {
		Question `xml:"question"`
	} `xml:"questions"`
	Type zfcs_publicDiscussionTypeEnum `xml:"type"` /* Тип обсуждений:
	LP - общественное обсуждение крупных закупок,
	OT - иное общественное обсуждение */
	Actual bool `xml:"actual"` /* Признак актуальности */
}

type zfcs_nsiPurchaseDocumentTypesType struct {
	PlacingWayCode string `xml:"placingWayCode"` /* Код подспособа размещения заказа (определения поставщика) */
	PlacingWayType string `xml:"placingWayType"` /* Код способа определения поставщика */
	PlacingWayName string `xml:"placingWayName"` /* Наименование способа определения поставщика */
	Actual         bool   `xml:"actual"`         /* Признак актуальности */
	DocumentTypes  struct {
		DocumentType `xml:"documentType"`
	} `xml:"documentTypes"`
}

type zfcs_nsiPurchasePreferencesType struct {
	Id        int64                   `xml:"id"`        /* Идентификатор преимущества (требования) */
	Name      string                  `xml:"name"`      /* Наименование преимущества (требования) */
	ShortName string                  `xml:"shortName"` /* Короткое наименование */
	Type      zfcs_preferenceTypeEnum `xml:"type"`      /* Тип преимущества (требования)

	P - Преимущество;
	F - Требование. */
	PrefEstimateApp bool    `xml:"prefEstimateApp"` /* Преимущество при оценке заявки */
	PrefValue       float64 `xml:"prefValue"`       /* Величина преимущества */
	PlacingWays     struct {
		Code string `xml:"code"`
	} `xml:"placingWays"`
	Actual         bool `xml:"actual"`         /* Признак актуальности */
	UseTenderPlans bool `xml:"useTenderPlans"` /* Используется в планах-графиках */
}

type zfcs_nsiPurchaseRejectReasonType struct {
	Id            int64                  `xml:"id"`            /* Идентификатор */
	Code          string                 `xml:"code"`          /* Код причины отказа */
	Reason        string                 `xml:"reason"`        /* Причина отказа */
	Actual        bool                   `xml:"actual"`        /* Признак актуальности */
	SubsystemType zfcs_subsystemTypeEnum `xml:"subsystemType"` /* К какому закону относится причина:

	FZ94 - 94ФЗ:
	FZ44 - 44ФЗ. */
}

type zfcs_nsiUserType struct {
	Login                string                    `xml:"login"`                /* Логин */
	Password             string                    `xml:"password"`             /* Хэш пароля */
	FirstName            string                    `xml:"firstName"`            /* Имя */
	MiddleName           string                    `xml:"middleName"`           /* Отчество */
	LastName             string                    `xml:"lastName"`             /* Фамилия */
	CodePhrase           string                    `xml:"codePhrase"`           /* Кодовая фраза */
	Position             string                    `xml:"position"`             /* Должность */
	Phone                zfcs_string               `xml:"phone"`                /* Телефон */
	Email                zfcs_string               `xml:"email"`                /* Адрес электронной почты */
	Organization         zfcs_organizationRef      `xml:"organization"`         /* Организация пользователя */
	CertificateSN        string                    `xml:"certificateSN"`        /* Серийный номер сертификата пользователя */
	CertificateMask      string                    `xml:"certificateMask"`      /* Отпечаток сертификата */
	EsIssuerDN           string                    `xml:"esIssuerDN"`           /* Полное наименование УЦ, выдавшего сертификат пользователя */
	EsIssuerSN           string                    `xml:"esIssuerSN"`           /* Cерийный номер сертификата УЦ, выдавшего сертификат пользователя */
	UserOrganizationRole zfcs_organizationRoleType `xml:"userOrganizationRole"` /* Полномочие организации, с которым связан пользователь организации:
	CU - Заказчик;
	RA - Уполномоченный орган;
	AI - Уполномоченное учреждение;
	SO - Специализированная организация;
	CO - Контрольный орган;
	SP - Служба Оператора сайта;
	FO - Финансовый орган;
	EO - Оператор электронной площадки;
	AA - Орган аудита;
	CA - Орган по регулированию контрактной системы в сфере закупок;
	NA - Орган, устанавливающий правила нормирования;
	DA - Орган, утверждающий требования к отдельным видам товаров, работ, услуг;
	BA - Банк;
	TA - Орган, разрабатывающий и утверждающий типовые контракты и типовые условия контрактов;
	OA - Организация, осуществляющая полномочия заказчика на осуществление закупок на основании договора (соглашения) в соответствии с частью 6 статьи 15 Закона № 44-ФЗ;
	CIA - Орган контроля соответствия информации об объемах финансового обеспечения и идентификационных кодах закупок;
	ICB - Орган внутреннего контроля;
	NP - Орган, уполномоченный на ведение реестра недобросовестных поставщиков;
	GR - Главный распорядитель бюджетных средств;
	OV - Орган государственной (исполнительной) власти. */
	Status zfcs_userStatusType `xml:"status"` /* Статус пользователя:
	A - Активный;
	B - Заблокирован;
	BO - Блокирование возможности размещения новых заказов.
	*/
	ETPPrivileges zfcs_etpPrivilege `xml:"ETPPrivileges"` /* Привилегии пользователя на ЭП */
}

type zfcs_organizationBudgetsType struct {
	RegNumber zfcs_string `xml:"regNumber"` /* Код Организации по Сводному Перечню Заказчиков (СПЗ) */
	Budgets   struct {
		Budget zfcs_nsiBudgetType `xml:"budget"`
	} `xml:"budgets"`
}

type zfcs_penalty_documentInfoList struct {
	DocumentInfo struct {
		DocumentDate string `xml:"documentDate"`
		DocumentName string `xml:"documentName"`
	} `xml:"documentInfo"`
}

type zfcs_planPositionChangeReasonRef struct {
	Id   int64  `xml:"id"`   /* Идентификатор */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_placementResultType struct {
	PurchaseNumber           zfcs_purchaseNumberType  `xml:"purchaseNumber"`           /* Номер закупки */
	ProtocolNumber           zfcs_documentNumberType  `xml:"protocolNumber"`           /* Номер протокола */
	LotNumber                uint64                   `xml:"lotNumber"`                /* Номер лота */
	FoundationProtocolNumber zfcs_documentNumberType  `xml:"foundationProtocolNumber"` /* Протокол – основание, являющийся итоговым протоколом для данного лота */
	VersionNumber            int                      `xml:"versionNumber"`            /* Номер редакции */
	CreateDate               string                   `xml:"createDate"`               /* Дата создания */
	ProcedurelFailed         bool                     `xml:"procedurelFailed"`         /* Признак несостоявшейся процедуры */
	AbandonedReason          zfcs_abandonedReasonType `xml:"abandonedReason"`          /* Основание признание процедуры несостоявшейся */
	Applications             struct {
		Application `xml:"application"`
	} `xml:"applications"`
	Result zfcs_placementResultEnum `xml:"result"` /* Результат определения поставщика:
	CC - Заключение контракта;
	RC - Отказ или уклонение от заключения контракта */
}

type zfcs_quickRefOrganizationType struct {
	RegNum          zfcs_string              `xml:"regNum"`          /* Код по СПЗ */
	ConsRegistryNum zfcs_consRegistryNumType `xml:"consRegistryNum"` /* Код по Сводному Реестру */
	FullName        zfcs_longTextType        `xml:"fullName"`        /* Полное наименование */
	FactAddress     zfcs_longTextType        `xml:"factAddress"`     /* Адрес местонахождения организации */
	INN             string                   `xml:"INN"`             /* ИНН организации */
	KPP             string                   `xml:"KPP"`             /* КПП организации */
	ContactEMail    string                   `xml:"contactEMail"`    /* Адрес электронной почты */
}

type zfcs_refusalFact struct {
	VoucherEntry zfcs_longTextType          `xml:"voucherEntry"` /* Реквизиты подтверждающих документов */
	Explanation  zfcs_longTextType          `xml:"explanation"`  /* Пояснения */
	Foundation   zfcs_refusalFactFoundation `xml:"foundation"`   /* Основание отказа */
}

type zfcs_refusalFactFoundation struct {
	Id   int64  `xml:"id"`   /* Идентификатор в справочнике */
	Name string `xml:"name"` /* Наименование */
}

type zfcs_deviationFactFoundation struct {
	Code string `xml:"code"` /* Код причины признания участника уклонившимся */
	Name string `xml:"name"` /* Наименование причины признания участника уклонившимся */
}

type zfcs_releasePurchaseDocumentationType struct {
	GrantStartDate string            `xml:"grantStartDate"` /* Дата и время начала предоставления (конкурсной/аукционной) документации по закупке */
	GrantPlace     zfcs_longTextType `xml:"grantPlace"`     /* Место предоставления (конкурсной/аукционной) документации по закупке */
	GrantOrder     zfcs_longTextType `xml:"grantOrder"`     /* Порядок предоставления (конкурсной/аукционной) документации по закупке */
	Languages      zfcs_longTextType `xml:"languages"`      /* Языки предоставления (конкурсной/аукционной) документации по закупке */
	GrantMeans     zfcs_longTextType `xml:"grantMeans"`     /* Способы предоставления (конкурсной/аукционной) документации по закупке */
	GrantEndDate   string            `xml:"grantEndDate"`   /* Дата и время окончания предоставления (конкурсной/аукционной) документации по закупке */
	PayCurrency    zfcs_currencyRef  `xml:"payCurrency"`    /* Валюта платежа за предоставление (конкурсной/аукционной) документации по закупке (Устарело) */
	PayInfo        `xml:"payInfo"`   /* Плата за предоставление (конкурсной/аукционной) документации по закупке. Отсутствие элемента в извещении означает, что плата за предоставление документации не установлена. */
}

type zfcs_requirementType struct {
	Code    int64             `xml:"code"`    /* Код требования */
	Name    zfcs_longTextType `xml:"name"`    /* Наименование требования */
	Content zfcs_longTextType `xml:"content"` /* Содержание требования */
}

type zfcs_nsiRMISType struct {
	RegNumber   zfcs_RMISNumType  `xml:"regNumber"`   /* Регистрационный номер */
	CreateDate  string            `xml:"createDate"`  /* Дата создания информации о системе */
	PublishDate string            `xml:"publishDate"` /* Дата размещения информации о системе */
	NameRMIS    zfcs_longTextType `xml:"nameRMIS"`    /* Наименование информационной системы */
	TypeRMIS    string            `xml:"typeRMIS"`    /* Тип информационной системы:

	IST_01-Региональная, муниципальная информационная система в сфере закупок;
	IST_02-Информационная система государственного внебюджетного фонда;
	IST_03-Информационная система оператора электронной площадки;
	IST_04-Информационная система контрольного органа в сфере закупок;
	IST_05-Информационная система органа внутреннего государственного (муниципального) финансового контроля;
	IST_06-Информационная система органа аудита в сфере закупок;
	IST_07-Иная информационная система в сфере закупок. */
	DescriptionRMIS      zfcs_longTextType      `xml:"descriptionRMIS"` /* Описание назначения информационной системы */
	StartDateRMIS        string                 `xml:"startDateRMIS"`   /* Дата начала постоянной эксплуатации информационной системы */
	RequisitesRMIS       zfcs_longTextType      `xml:"requisitesRMIS"`  /* Реквизиты документа (нормативного правового акта) о вводе информационной системы в постоянную эксплуатацию */
	UrlRMIS              zfcs_longTextType      `xml:"urlRMIS"`         /* Адрес (URL) информационной системы в информационно-коммуникационной сети "Интернет" */
	PPO                  zfcs_PPORef            `xml:"PPO"`             /* Публично-правовое образование, на территории которого создана информационная система */
	TOFK                 zfcs_TOFKRef           `xml:"TOFK"`            /* Публично-правовое образование, на территории которого создана информационная система */
	FullName             zfcs_longTextType      `xml:"fullName"`        /* Полное наименование организации, являющейся оператором информационной системы */
	CodeSPZ              zfcs_string            `xml:"codeSPZ"`         /* Код организации, являющейся оператором информационной системы, по Сводному перечню заказчиков (СПЗ) */
	CodeSvodReestr       zfcs_svodReestrNumType `xml:"codeSvodReestr"`  /* Код организации, являющейся оператором информационной системы, по Сводному реестру */
	OGRN                 zfcs_ogrnType          `xml:"OGRN"`            /* Код ОГРН организации, являющейся оператором информационной системы */
	INN                  zfcs_innType           `xml:"INN"`             /* Код ИНН организации, являющейся оператором информационной системы */
	KPP                  zfcs_kppType           `xml:"KPP"`             /* Код КПП организации, являющейся оператором информационной системы */
	AuthorizedPersonInfo struct {
		LastName   zfcs_longTextType `xml:"lastName"`
		FirstName  zfcs_longTextType `xml:"firstName"`
		MiddleName zfcs_longTextType `xml:"middleName"`
		Position   zfcs_longTextType `xml:"position"`
		Address    zfcs_longTextType `xml:"address"`
		Phone      zfcs_longTextType `xml:"phone"`
		Email      zfcs_longTextType `xml:"email"`
	} `xml:"authorizedPersonInfo"`
	CertificateInfo struct {
		Content string `xml:"content"`
	} `xml:"certificateInfo"`
	Actual      bool                    `xml:"actual"`      /* Признак актуальности */
	Attachments zfcs_attachmentListType `xml:"attachments"` /* Информация о прикрепленных документах информационной системы */
}

type zfcs_nsiBusinessControlType struct {
	GUID      string                               `xml:"GUID"`      /* Глобально-уникальный идентификатор записи */
	Code      string                               `xml:"code"`      /* Бизнес-идентификатор записи */
	Order     uint64                               `xml:"order"`     /* Порядок следования контроля */
	Name      zfcs_longTextType                    `xml:"name"`      /* Наименование контроля */
	Actual    bool                                 `xml:"actual"`    /* Актуальность сведений */
	SubSystem zfcs_nsiBusinessControlSubSystemEnum `xml:"subSystem"` /* Подсистема контроля:
	PRIZ - Подсистема размещения информации о закупках (ПРИЗ);
	RK - Реестр контрактов (РК);
	RBG - Реестр банковских гарантий (РБГ);
	RPZ - Реестр планов закупок (РПЗ);
	RPG - Реестр планов-графиков (РПГ);
	OZ - Отчеты заказчиков (ОЗ);
	RDI - Реестр дополнительной нформации (РДИ);
	OOKZ - Общественное обсуждение крупных закупок (ООКЗ);
	BTK - Библиотека типовых контрактов (БТК);
	RPNZ - Реестр правил нормирования закупок (РПНЗ);
	ZC - Запрос цен (ЗЦ);
	RZH - Реестр жалоб (РЖ);
	RRK - Реестр результатов контроля (РРК);
	RPP - Реестр плановых проверок (РПП);
	RVP - Реестр внеплановых проверок (РВП);
	RRA - Реестр результатов аудита.
	*/
	Document zfcs_nsiBusinessControlDocumentEnum `xml:"document"` /* Документ контроля:

	contract	Информация о заключенном контракте
	contractCancel	Информация об анулировании контракта
	contractProcedure	Информация об исполнении (расторжении) контракта
	contractProcedureCancel	Сведения об отмене информации об исполнении (расторжении) контракта
	bankGuarantee	Информация о выданной банковской гарантии
	bankGuaranteeInvalid	Информация о недействительности сведений о банковской гарантии
	bankGuaranteeRefusal	Сведения об отказе заказчика в принятии банковской гарантии
	bankGuaranteeRefusalInvalid	Сведения о недействительности отказа заказчика в принятии банковской гарантии
	bankGuaranteeTermination	Информация о прекращении обязательств поставщика по банковской гарантии; внесение изменений
	bankGuaranteeTerminationInvalid	Сведения о недействительности информации о прекращении обязательств поставщика по банковской гарантии
	bankGuaranteeReturn	Информация о возвращении банковской гарантии или об освобождении от обязательств по банковской гарантии; внесение изменений
	bankGuaranteeReturnInvalid	Сведения о недействительности информации о возвращении банковской гарантии или об освобождении от обязательств по банковской гарантии
	complaint	Информация по жалобе
	complaintGroup	Информация по групповой жалобе
	complaintCancel	Информация об отзыве жалобы
	tenderSuspension	Информация о приостановке определения поставщика
	checkResult	Результат контроля
	checkResultCancel	Информация об отмене результата контроля
	unfairSupplier	Информация о недобросовестном поставщике
	unplannedCheck	Внеплановая проверка
	unplannedCheckCancel	Информация об отмене проведения проверки
	checkPlan	План проверок
	tenderPlan	План-график размещения заказов
	tenderPlanChange	Изменение плана-графика размещения заказов
	tenderPlanUnstructured	План-график размещения заказов в неструктурированной форме
	tenderPlanCancel	Аннулированиe плана-графика
	sketchPlan	План закупок в неструктурированной форме
	sketchPlanExecution	Сведения об исполнении плана закупок
	fcsNotificationEF	Извещение о проведении ЭА (электронный аукцион)
	fcsProtocolEF1	Протокол рассмотрения заявок на участие в электронном аукционе
	fcsProtocolEF2	Протокол проведения электронного аукциона
	fcsProtocolEF3	Протокол подведения итогов электронного аукциона
	fcsProtocolEFSingleApp	Протокол рассмотрения единственной заявки на участие в электронном аукционе
	fcsProtocolEFSinglePart	Протокол рассмотрения заявки единственного участника электронного аукциона
	fcsProtocolEFInvalidation	Протокол о признании электронного аукциона несостоявшимся
	fcsNotificationEP	Извещение о проведении закупки у ЕП (единственного поставщика)
	fcsNotificationOK	Извещение о проведении OK (открытый конкурс)
	fcsPurchaseProlongationOK	Уведомление о продлении срока рассмотрения и оценки заявок ОК
	fcsProtocolOK1	Протокол вскрытия конвертов и открытия доступа к электронным документам заявок участников в ОК
	fcsProtocolOK2	Протокол рассмотрения и оценки заявок на участие в конкурсе в ОК
	fcsProtocolOKSingleApp	Протокол рассмотрения единственной заявки в ОК; внесение изменений
	fcsNotificationOKD	Извещение о проведении OK-Д (двухэтапный конкурс)
	fcsProtocolOKD1	Протокол вскрытия конвертов и открытия доступа к электронным документам первоначальных заявок в ОК-Д
	fcsProtocolOKD2	Протокол предквалификационного отбора в ОК-Д
	fcsProtocolOKD3	Протокол первого этапа в ОК-Д
	fcsProtocolOKD4	Протокол вскрытия конвертов и открытия доступа к электронным документам окончательных заявок в ОК-Д
	fcsProtocolOKD5	Протокол рассмотрения и оценки заявок на участие в конкурсе в ОК-Д
	fcsProtocolOKDSingleApp	Протокол рассмотрения единственной заявки в ОК-Д
	fcsNotificationOKOU	Извещение о проведении OK-ОУ (конкурс с ограниченным участием)
	fcsProtocolOKOU1	Протокол вскрытия конвертов и открытия доступа к электронным документам заявок участников в ОК-ОУ
	fcsProtocolOKOU2	Протокол предквалификационного отбора в ОК-ОУ
	fcsProtocolOKOU3	Протокол рассмотрения и оценки заявок на участие в конкурсе в ОК-ОУ
	fcsProtocolOKOUSingleApp	Протокол рассмотрения единственной заявки в ОК-ОУ
	fcsNotificationPO	Извещение о проведении ПО (предварительный отбор)
	fcsProtocolPO	Протокол предварительного отбора в ПО
	fcsNotificationZakA	Извещение о проведении ЗакА (закрытый аукцион)
	fcsNotificationZakK	Извещение о проведении ЗакK (закрытый конкурс)
	fcsNotificationZakKD	Извещение о проведении ЗакK-Д (закрытый двухэтапный конкурс)
	fcsNotificationZakKOU	Извещение о проведении ЗакK-ОУ (закрытый конкурс с ограниченным участием)
	fcsNotificationZKBI	Общая информация об объекте закупки и протокол рассмотрения и оценки заявок на участие в ЗК-БИ (запрос котировок без извещения)
	fcsProtocolZKBI	Общая информация об объекте закупки и структурированный протокол рассмотрения и оценки заявок на участие в ЗК-БИ (запрос котировок без извещения);
	fcsNotificationZK	Извещение о проведении ЗK (запрос котировок)
	fcsPurchaseProlongationZK	Извещение о продлении срока подачи котировочных заявок в ЗК
	fcsProtocolZK	Протокол рассмотрения и оценки заявок в ЗК
	fcsProtocolZKAfterProlong	Протокол рассмотрения и оценки заявок по результатам продления срока подачи заявок в ЗК
	fcsNotificationZP	Извещение о проведении ЗП (запроса предложений)
	fcsProtocolZPExtract	Протокол выписки из протокола проведения запроса предложений в ЗП
	fcsProtocolZP	Протокол проведения запроса предложений в ЗП
	fcsProtocolZPFinal	Итоговый протокол в ЗП
	fcsProtocolCancel	Информация об отмене протокола
	fcsNotificationLotCancel	Извещение об отмене определения поставщика (подрядчика, исполнителя) в части лота
	fcsNotificationLotChange	Внесение изменений в извещение в части лота
	fcsNotificationCancel	Извещение об отмене определения поставщика (подрядчика, исполнителя)
	fcsNotificationCancelFailure	Отмена извещения об отмене определения поставщика (подрядчика, исполнителя) (в части лота)
	fcsNotificationOrgChange	Уведомление об изменении организации, осуществляющей закупку
	fcsProtocolEvasion	Протокол отказа от заключения контракта
	protocolDeviation	Протокол признания участника уклонившимся от заключения контракта
	fcsNotificationEFDateChange	Уведомление об изменении даты и времени проведения электронного аукциона
	fcsContractSign	Информация о подписании государственного/муниципального контракта на ЭП по 44-ФЗ
	fcsClarificationRequest	Запрос ЭП о даче разъяснений положений документации по 44-ФЗ
	fcsClarification	Разъяснения положений документации
	fcsTimeEF	Время проведения электронного аукциона
	masterData	Справочные данные
	purchaseDocument	Информация о документе закупки
	purchaseDocumentCancel	Информация об отмене документа закупки
	fcsPlacementResult	Результат проведения процедуры определения поставщика
	publicDiscussionLargePurchase	Информация об обязательном общественном обсуждения крупной закупки
	publicDiscussionAnwser	Ответ на комментарий обязательного общественного обсуждения крупной закупки
	publicDiscussionProtocol	Протокол первого/второго этапа обязательного общественного обсуждения крупной закупки
	customerReportContractExecution	Отчет об исполнении контракта (результатах отдельного этапа исполнения контракта)
	customerReportContractExecutionInvalid	Информация о недействительности сведений отчета об исполнении контракта (результатах отдельного этапа исполнения контракта)
	customerReportSmallScaleBusiness	Отчет об объеме закупок у СМП (субъектов малого предпринимательства), СОНО (социально ориентированных некоммерческих организаций)
	customerReportSmallScaleBusinessInvalid	Информация о недействительности сведений отчета об объеме закупок у СМП (субъектов малого предпринимательства), СОНО (социально ориентированных некоммерческих организаций)
	customerReportSingleContractor	Отчет с обоснованием закупки у единственного поставщика (подрядчика, исполнителя)
	customerReportSingleContractorInvalid	Информация о недействительности сведений отчета с обоснованием закупки у единственного поставщика (подрядчика, исполнителя)
	customerReportBigProjectMonitoring	Отчет по мониторингу реализации крупных проектов с государственным участием
	customerReportBigProjectMonitoringInvalid	Информация о недействительности сведений отчета по мониторингу реализации крупных проектов с государственным участием
	addInfo	Дополнительная информация о закупках, контрактах
	addInfoInvalid	Информация о недействительности дополнительной информации о закупках, контрактах
	requestForQuotation	Версия запроса цен товаров, работ, услуг предусмотренных частью 5 статьи 22 закона №44-ФЗ
	requestForQuotationCancel	Отмена запроса цен товаров, работ, услуг
	auditResult	Обобщенная информация о результатах деятельности органа аудита в сфере закупок
	eventPlan	План мероприятий
	eventResult	Результат контроля
	eventResultCancel	Информация об отмене результата контроля
	unplannedEvent	Информация по внеплановому контрольному мероприятию
	unplannedEventCancel	Информация об отмене проведения внепланового контрольного мероприятия
	*/
	Action zfcs_nsiBusinessControlActionEnum `xml:"action"` /* Действие с объектом, на которое срабатывает контроль:
	Add - Сохранение проекта документа;
	Edit - Рeдактирование проекта документа;
	Delete - Удаление проекта документа;
	Publish - Размещение (публикация) документа */
	Level zfcs_nsiBusinessControlLevelEnum `xml:"level"` /* Уровень контроля:
	0 - Ошибка. Не допускает сохранение;
	1 - Ошибка. Не допускает размещение(публикацию) ;
	2 - Ошибка. Позволяет размещение(публикацию) ;
	3 - Предупреждение. Позволяет размещение(публикацию) ; */
	Description zfcs_longTextType `xml:"description"` /* Описание контроля */
}

type zfcs_nsiBusinessControlsType struct {
	NsiBusinessControl zfcs_nsiBusinessControlType `xml:"nsiBusinessControl"` /* Бизнес-контроль */
}

type zfcs_TOFKRef struct {
	Code zfcs_TOFKCodeType `xml:"code"` /* Код ТОФК */
	Name zfcs_longTextType `xml:"name"` /* Наименование ТОФК */
}

type zfcs_sketchPlanExecutionType struct {
	CommonInfo struct {
		Id             int64                     `xml:"id"`
		ExternalId     zfcs_externalIdType       `xml:"externalId"`
		PlanNumber     zfcs_sketchPlanNumberType `xml:"planNumber"`
		Year           zfcs_yearType             `xml:"year"`
		PeriodYearFrom zfcs_yearType             `xml:"periodYearFrom"`
		PeriodYearTo   zfcs_yearType             `xml:"periodYearTo"`
		Name           string                    `xml:"name"`
		VersionNumber  int                       `xml:"versionNumber"`
	} `xml:"commonInfo"`
	CustomerInfo  zfcs_quickRefOrganizationType `xml:"customerInfo"`  /* Сведения о заказчике */
	CreateDate    string                        `xml:"createDate"`    /* Дата создания исполнения плана закупок */
	VersionNumber int                           `xml:"versionNumber"` /* Номер версии исполнения плана закупок */
	PublishDate   string                        `xml:"publishDate"`   /* Дата публикации исполнения плана закупок */
	Description   zfcs_longTextType             `xml:"description"`   /* Описание сведений об исполнении плана закупок */
	Attachments   zfcs_attachmentListType       `xml:"attachments"`   /* Информация о прикрепленных документах */
	PrintForm     zfcs_printFormType            `xml:"printForm"`     /* Печатная форма исполгнения плана закупок */
	ExtPrintForm  zfcs_extPrintFormType         `xml:"extPrintForm"`  /* Электронный документ, полученный из внешней системы */
}

type zfcs_sketchPlanType struct {
	CommonInfo struct {
		Id             int64                     `xml:"id"`
		ExternalId     zfcs_externalIdType       `xml:"externalId"`
		PlanNumber     zfcs_sketchPlanNumberType `xml:"planNumber"`
		CreateDate     string                    `xml:"createDate"`
		Year           zfcs_yearType             `xml:"year"`
		PeriodYearFrom zfcs_yearType             `xml:"periodYearFrom"`
		PeriodYearTo   zfcs_yearType             `xml:"periodYearTo"`
		Name           zfcs_longTextType         `xml:"name"`
		Owner          zfcs_organizationRef      `xml:"owner"`
		PublishDate    string                    `xml:"publishDate"`
		VersionNumber  int                       `xml:"versionNumber"`
		ConfirmDate    string                    `xml:"confirmDate"`
	} `xml:"commonInfo"`
	CustomerInfo zfcs_quickRefOrganizationType `xml:"customerInfo"` /* Сведения о заказчике */
	Attachments  zfcs_attachmentListType       `xml:"attachments"`  /* Информация о прикрепленных документах */
	PrintForm    zfcs_printFormType            `xml:"printForm"`    /* Печатная форма плана закупок */
	ExtPrintForm zfcs_extPrintFormType         `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_tenderPlanType struct {
	CommonInfo   zfcs_tenderPlanCommonInfoType `xml:"commonInfo"` /* Общая информация плана-графика */
	CustomerInfo struct {
		Customer zfcs_purchaseOrganizationType `xml:"customer"`
		OKTMO    zfcs_OKTMORef                 `xml:"OKTMO"`
	} `xml:"customerInfo"`
	ResponsibleContactInfo `xml:"responsibleContactInfo"` /* Сведения об исполнителе (ответственном за формирование плана-графика) */
	ProvidedPurchases      struct {
		Positions      `xml:"positions"`
		FinalPositions zfcs_tenderPlanFinalPositionsType `xml:"finalPositions"`
	} `xml:"providedPurchases"`
	PrintForm    zfcs_printFormType    `xml:"printForm"`    /* Печатная форма плана-графика */
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_tenderPlanCancelType struct {
	Id            int64                 `xml:"id"`         /* Идентификатор аннулированного плана-графика на ЕИС */
	PlanNumber    string                `xml:"planNumber"` /* Реестровый номер плана-графика */
	CustomerInfo  `xml:"customerInfo"`  /* Сведения о заказчике (для печатной формы) */
	Year          zfcs_yearType         `xml:"year"`          /* Год плана-графика */
	Description   zfcs_longTextType     `xml:"description"`   /* Основание для аннулирования плана-графика */
	VersionNumber int                   `xml:"versionNumber"` /* Номер версии плана-графика */
	CancelDate    string                `xml:"cancelDate"`    /* Дата и время аннулирования (публикации аннулированного плана-графика) */
	PrintForm     zfcs_printFormType    `xml:"printForm"`     /* Печатная форма сведений об аннулировании плана-графика */
	ExtPrintForm  zfcs_extPrintFormType `xml:"extPrintForm"`  /* Электронный документ, полученный из внешней системы */
}

type zfcs_tenderPlanUnstructuredType struct {
	CommonInfo          zfcs_tenderPlanCommonInfoType `xml:"commonInfo"`          /* Общая информация плана-графика */
	PublicDiscussionNum zfcs_publicDiscussionNumType  `xml:"publicDiscussionNum"` /* Реестровый номер общественного обсуждения */
	CustomerInfo        struct {
		Customer zfcs_organizationRef `xml:"customer"`
		OKTMO    zfcs_OKTMORef        `xml:"OKTMO"`
	} `xml:"customerInfo"`
	ResponsibleContactInfo `xml:"responsibleContactInfo"` /* Сведения об исполнителе (ответственном за формирование плана-графика) */
	ProvidedNotPurchases   bool                           `xml:"providedNotPurchases"` /* Закупки не предусмотрены */
	Attachments            zfcs_attachmentListType        `xml:"attachments"`          /* Информация о прикрепленных документах */
	PrintForm              zfcs_printFormType             `xml:"printForm"`            /* Печатная форма плана-графика */
}

type zfcs_tenderSuspensionType struct {
	ComplaintNumber zfcs_complaintNumberType        `xml:"complaintNumber"` /* Номер жалобы */
	RegDate         string                          `xml:"regDate"`         /* Дата решения о приостановке определения поставщика */
	PublishDate     string                          `xml:"publishDate"`     /* Дата публикации решения  */
	Action          zfcs_tenderSuspensionActionType `xml:"action"`          /* Тип действия
	SUSPENSION - приостановка
	REOPENING - возобновление */
	KO          zfcs_organizationRef `xml:"KO"` /* Контролирующий орган, принявший решение о приостановке определения поставщика */
	TendersInfo `xml:"tendersInfo"`  /* Информация о заказе/закупке */
}

type zfcs_tendePlanInfoType struct {
	PlanNumber     zfcs_tenderPlanNumberType         `xml:"planNumber"`     /* Реестровый номер плана-графика */
	PositionNumber zfcs_tenderPlanPositionNumberType `xml:"positionNumber"` /* Номер позиции в плане-графике */
}

type zfcs_tenderPlan_ContextType struct {
}

type zfcs_tenderPlanChangeType struct {
	Id                     int64                     `xml:"id"`                     /* Идентификатор плана-графика на ЕИС */
	PlanNumber             zfcs_tenderPlanNumberType `xml:"planNumber"`             /* Реестровый номер плана-графика */
	VersionNumber          int                       `xml:"versionNumber"`          /* Номер версии плана-графика */
	Description            string                    `xml:"description"`            /* Описание плана-графика */
	ConfirmDate            string                    `xml:"confirmDate"`            /* Дата утверждения плана-графика */
	ResponsibleContactInfo zfcs_userTenderPlanType   `xml:"responsibleContactInfo"` /* Сведения об исполнителе (ответственном за формирование плана-графика)  */
	ProvidedPurchases      struct {
		Positions      `xml:"positions"`
		FinalPositions zfcs_tenderPlanFinalPositionsType `xml:"finalPositions"`
	} `xml:"providedPurchases"`
	ExtPrintForm zfcs_extPrintFormType `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_tenderPlanCommonInfoType struct {
	Id            int64                     `xml:"id"`            /* Идентификатор плана-графика на ЕИС */
	ExternalId    zfcs_externalIdType       `xml:"externalId"`    /* Внешний идентификатор документа */
	PlanNumber    zfcs_tenderPlanNumberType `xml:"planNumber"`    /* Реестровый номер плана-графика */
	Year          zfcs_yearType             `xml:"year"`          /* Год плана-графика */
	VersionNumber int                       `xml:"versionNumber"` /* Номер версии плана-графика */
	Owner         `xml:"owner"`             /* Организация владельца версии плана-графика */
	CreateDate    string                    `xml:"createDate"`  /* Дата создания плана-графика */
	Description   zfcs_longTextType         `xml:"description"` /* Описание плана-графика */
	ConfirmDate   string                    `xml:"confirmDate"` /* Дата утверждения плана-графика */
	PublishDate   string                    `xml:"publishDate"` /* Дата публикации плана-графика */
}

type zfcs_tenderPlanFinalPositionsType struct {
	Purchase83 struct {
		TeachingService `xml:"teachingService"`
		GuideService    `xml:"guideService"`
		Medicine        `xml:"medicine"`
	} `xml:"purchase83"`
	Purchase83st544 struct {
		Medicine `xml:"medicine"`
	} `xml:"purchase83st544"`
	Purchase93 struct {
		PurchaseAmountLess100      `xml:"purchaseAmountLess100"`
		PurchaseAmountLess400      `xml:"purchaseAmountLess400"`
		MaintenanceRepairService   `xml:"maintenanceRepairService"`
		BusinessTripService        `xml:"businessTripService"`
		TeachingService            `xml:"teachingService"`
		GuideService               `xml:"guideService"`
		CollectionStatisticService `xml:"collectionStatisticService"`
		AccessDBService            `xml:"accessDBService"`
	} `xml:"purchase93"`
	OutcomeIndicators struct {
		SumPushaseSingleSupplier4 zfcs_moneyPositiveType `xml:"sumPushaseSingleSupplier4"`
		SumPushaseSingleSupplier5 zfcs_moneyPositiveType `xml:"sumPushaseSingleSupplier5"`
		SumPushaseSmallBusiness   zfcs_moneyPositiveType `xml:"sumPushaseSmallBusiness"`
		SumPushaseRequest         zfcs_moneyPositiveType `xml:"sumPushaseRequest"`
		SumContractMaxPrice       zfcs_moneyPositiveType `xml:"sumContractMaxPrice"`
		SumPaymentsTotal          zfcs_moneyPositiveType `xml:"sumPaymentsTotal"`
	} `xml:"outcomeIndicators"`
}

type zfcs_tenderPlanPositionType struct {
	CommonInfo struct {
		PositionNumber        zfcs_tenderPlanPositionNumberType `xml:"positionNumber"`
		ExtNumber             int                               `xml:"extNumber"`
		OrderNumber           string                            `xml:"orderNumber"`
		OKVEDs                `xml:"OKVEDs"`
		ContractSubjectName   zfcs_longTextType      `xml:"contractSubjectName"`
		LegalActRequisites    zfcs_longTextType      `xml:"legalActRequisites"`
		ContractMaxPrice      zfcs_moneyPositiveType `xml:"contractMaxPrice"`
		Payments              zfcs_moneyPositiveType `xml:"payments"`
		ContractPriceFeatures zfcs_longTextType      `xml:"contractPriceFeatures"`
		ContractCurrency      zfcs_currencyRef       `xml:"contractCurrency"`
		PlacingWay            zfcs_placingWayType    `xml:"placingWay"`
		Features111           bool                   `xml:"features111"`
		JointBiddingInfo      `xml:"jointBiddingInfo"`
		PositionModification  `xml:"positionModification"`
		PositionPublishDate   string `xml:"positionPublishDate"`
	} `xml:"commonInfo"`
	Products struct {
		Product `xml:"product"`
	} `xml:"products"`
	PurchaseConditions struct {
		PurchaseFinCondition   `xml:"purchaseFinCondition"`
		ContractFinCondition   `xml:"contractFinCondition"`
		Advance                zfcs_longTextType `xml:"advance"`
		PurchaseGraph          `xml:"purchaseGraph"`
		Prohibitions           zfcs_longTextType `xml:"prohibitions"`
		PreferensesRequirement `xml:"preferensesRequirement"`
	} `xml:"purchaseConditions"`
}

type zfcs_tenderPlanTotalPositionKBKsType struct {
	KBK struct {
		Code   zfcs_kbkType           `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KBK"`
}

type zfcs_tenderPlanTotalPositionKOSGUsType struct {
	KOSGU struct {
		Code   zfcs_kosguType         `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KOSGU"`
}

type zfcs_tenderPlanTotalPositionKBK2016sType struct {
	KBK2016 struct {
		Code   zfcs_kbk2016Type       `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KBK2016"`
}

type zfcs_tenderPlanTotalPositionKVRsType struct {
	KVR struct {
		Code   zfcs_KVRCodeType       `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KVR"`
}

type zfcs_userTenderPlanType struct {
	LastName   string `xml:"lastName"`   /* Фамилия */
	FirstName  string `xml:"firstName"`  /* Имя */
	MiddleName string `xml:"middleName"` /* Отчество */
	Phone      string `xml:"phone"`      /* Телефон */
	Fax        string `xml:"fax"`        /* Факс  */
	Email      string `xml:"email"`      /* Адрес электронной почты */
}

type zfcs_tenderPlanPositionKBKsType struct {
	KBK struct {
		Code   zfcs_kbkType           `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KBK"`
}

type zfcs_tenderPlanPositionKOSGUsType struct {
	KOSGU struct {
		Code   zfcs_kosguType         `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KOSGU"`
}

type zfcs_tenderPlanPositionKBK2016sType struct {
	KBK2016 struct {
		Code   zfcs_kbk2016Type       `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KBK2016"`
}

type zfcs_tenderPlanPositionKVRsType struct {
	KVR struct {
		Code   zfcs_KVRCodeType       `xml:"code"`
		Amount zfcs_moneyPositiveType `xml:"amount"`
	} `xml:"KVR"`
}

type zfcs_PositionKBKsYearsType struct {
	KBK struct {
		Code      zfcs_kbkType `xml:"code"`
		YearsList `xml:"yearsList"`
	} `xml:"KBK"`
}

type zfcs_PositionKOSGUsYearsType struct {
	KOSGU struct {
		Code      zfcs_kosguType `xml:"code"`
		YearsList `xml:"yearsList"`
	} `xml:"KOSGU"`
}

type zfcs_PositionKBK2016sYearsType struct {
	KBK2016 struct {
		Code      zfcs_kbk2016Type `xml:"code"`
		YearsList `xml:"yearsList"`
	} `xml:"KBK2016"`
}

type zfcs_PositionKVRsYearsType struct {
	KVR struct {
		Code      zfcs_KVRCodeType `xml:"code"`
		YearsList `xml:"yearsList"`
	} `xml:"KVR"`
}

type zfcs_purchasePlanType struct {
	CommonInfo struct {
		Id            int64                     `xml:"id"`
		ExternalId    zfcs_externalIdType       `xml:"externalId"`
		PlanNumber    zfcs_sketchPlanNumberType `xml:"planNumber"`
		PlanPeriod    `xml:"planPeriod"`
		VersionNumber int `xml:"versionNumber"`
		Owner         `xml:"owner"`
		CreateDate    string                        `xml:"createDate"`
		ConfirmDate   string                        `xml:"confirmDate"`
		Base          zfcs_purchasePlanBasementEnum `xml:"base"`
		PublishDate   string                        `xml:"publishDate"`
	} `xml:"commonInfo"`
	CustomerInfo struct {
		Customer               `xml:"customer"`
		ResponsibleContactInfo zfcs_contactPersonType `xml:"responsibleContactInfo"`
		ConfirmContactInfo     `xml:"confirmContactInfo"`
	} `xml:"customerInfo"`
	LocalInfo         zfcs_purchasePlanAddInfoType `xml:"localInfo"` /* Дополнительные сведения о плане закупок, установленные субъектом или муниципальным образованием Российской Федерации */
	ProvidedPurchases struct {
		Positions             `xml:"positions"`
		SpecialPurchases      `xml:"specialPurchases"`
		Finances              `xml:"finances"`
		TotalContractFinances zfcs_financeResourcesType `xml:"totalContractFinances"`
		TotalPurchaseFinances zfcs_financeResourcesType `xml:"totalPurchaseFinances"`
	} `xml:"providedPurchases"`
	Attachments  zfcs_attachmentListType `xml:"attachments"`  /* Информация о прикрепленных документах */
	PrintForm    zfcs_printFormType      `xml:"printForm"`    /* Печатная форма плана закупок */
	ExtPrintForm zfcs_extPrintFormType   `xml:"extPrintForm"` /* Электронный документ, полученный из внешней системы */
}

type zfcs_purchasePlanPositionType struct {
	CommonInfo struct {
		PositionNumber       zfcs_tenderPlanPositionNumberType `xml:"positionNumber"`
		IKZ                  zfcs_ikzCodeType                  `xml:"IKZ"`
		OKPD                 zfcs_OKPDRef                      `xml:"OKPD"`
		PurchaseObjectInfo   zfcs_longTextType                 `xml:"purchaseObjectInfo"`
		Description          zfcs_longTextType                 `xml:"description"`
		Complexity           bool                              `xml:"complexity"`
		ComplexityInfo       zfcs_longTextType                 `xml:"complexityInfo"`
		PublicDiscussion     bool                              `xml:"publicDiscussion"`
		Execution            `xml:"execution"`
		Basement             zfcs_longTextMinType `xml:"basement"`
		Purpose              `xml:"purpose"`
		PositionCanceled     bool `xml:"positionCanceled"`
		PositionModification `xml:"positionModification"`
	} `xml:"commonInfo"`
	FinanceInfo struct {
		PublishYear  zfcs_yearType                           `xml:"publishYear"`
		Finances     zfcs_purchasePlanPositionFinancingsType `xml:"finances"`
		YearFinances zfcs_financeResourcesType               `xml:"yearFinances"`
	} `xml:"financeInfo"`
	LocalInfo zfcs_purchasePlanAddInfoType `xml:"localInfo"` /* Дополнительные сведения о позиции плана закупок, установленные субъектом или муниципальным образованием Российской Федерации */
}

type zfcs_purchasePlanOrganizationType struct {
	RegNum          zfcs_string              `xml:"regNum"`          /* Код по СПЗ */
	ConsRegistryNum zfcs_consRegistryNumType `xml:"consRegistryNum"` /* Код по Сводному Реестру */
	FullName        zfcs_longTextMinType     `xml:"fullName"`        /* Полное наименование */
	FactAddress     zfcs_longTextMinType     `xml:"factAddress"`     /* Адрес местонахождения организации */
	Phone           string                   `xml:"phone"`           /* Телефон */
	Email           string                   `xml:"email"`           /* Адрес электронной почты */
}

type zfcs_purchasePlanPositionFinancingsType struct {
}

type zfcs_financeResourcesType struct {
	CurrentYear zfcs_moneyPlanType `xml:"currentYear"` /* Объем финансового обеспечения на текущий финансовый год */
	FirstYear   zfcs_moneyPlanType `xml:"firstYear"`   /* Объем финансового обеспечения на первый плановый финансовый год */
	SecondYear  zfcs_moneyPlanType `xml:"secondYear"`  /* Объем финансового обеспечения на второй плановый финансовый год */
	SubsecYears zfcs_moneyPlanType `xml:"subsecYears"` /* Объем финансового обеспечения на последующие годы */
	Total       zfcs_moneyPlanType `xml:"total"`       /* Общий объем финансового обеспечения */
}

type zfcs_specialPurchaseRef struct {
	Code string               `xml:"code"` /* Код типа */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование типа особых закупок в планах закупок */
}

type zfcs_purchasePlanPositionChangeReasonRef struct {
	Code string               `xml:"code"` /* Код основания */
	Name zfcs_longTextMinType `xml:"name"` /* Наименование */
}

type zfcs_purchasePlanAddInfoType struct {
	AddInfo    zfcs_longTextType   `xml:"addInfo"`    /* Дополнительные сведения */
	Attachment zfcs_attachmentType `xml:"attachment"` /* Прикрепленный документ */
}

type zfcs_unfairSupplierType struct {
	RegistryNum zfcs_documentNumberType `xml:"registryNum"` /* Реестровый номер */
	PublishDate string                  `xml:"publishDate"` /* Дата публикации */
	ApproveDate string                  `xml:"approveDate"` /* Дата включения сведений в реестр */
	State       string                  `xml:"state"`       /* Статус информации
	REVIEW_CANCELED - Отклонена (Отказ во включении в РНП);
	PUBLISHED - Опубликована;
	ON_EXCLUSION - Заявка на исключение сведений;
	TEMP_EXCLUDED - Информация исключена из РНП на время судебного разбирательства;
	ARCHIVED - Информация исключена из РНП. Архив; */
	PublishOrg   zfcs_organizationRef `xml:"publishOrg"`   /* Уполномоченный орган, осуществивший включение информации в реестр */
	CreateReason zfcs_longTextType    `xml:"createReason"` /* Причина для внесения в реестр: WINNER_DEVIATION - уклонение победителя от заключения контракта
	ONE_WINNER_DEVIATION - уклонение единственного участника от заключения контракта
	PARTICIPANT_DEVIATION_IF_WINNER_DEVIATION - уклонение участника размещения заказа от заключения контракта в случае, когда победитель уклонился от заключения контракта
	CANCEL_CONTRACT - расторжение контракта */
	ApproveReason  zfcs_longTextType             `xml:"approveReason"` /* Основание для включения в РНП */
	Customer       zfcs_purchaseOrganizationType `xml:"customer"`      /* Информация о заказчике, подавшем заявку на включение реестр */
	UnfairSupplier struct {
		FullName zfcs_longTextType        `xml:"fullName"`
		Type     zfcs_participantTypeEnum `xml:"type"`
		FirmName zfcs_longTextType        `xml:"firmName"`
		Inn      string                   `xml:"inn"`
		Kpp      zfcs_kppType             `xml:"kpp"`
		Place    `xml:"place"`
		Founders `xml:"founders"`
	} `xml:"unfairSupplier"`
	Purchase struct {
		PurchaseNumber     zfcs_purchaseNumberType `xml:"purchaseNumber"`
		PurchaseObjectInfo zfcs_longTextType       `xml:"purchaseObjectInfo"`
		PlacingWayName     string                  `xml:"placingWayName"`
		ProtocolDate       string                  `xml:"protocolDate"`
		LotNumber          int                     `xml:"lotNumber"`
		Document           `xml:"document"`
	} `xml:"purchase"`
	NotOosPurchase bool `xml:"notOosPurchase"` /* Извещение о проведении закупки не размещено на официальном сайте */
	Contract       struct {
		RegNum      zfcs_contract_regNumType `xml:"regNum"`
		ProductInfo zfcs_longTextType        `xml:"productInfo"`
		OKPD        zfcs_OKPDRef             `xml:"OKPD"`
		Currency    zfcs_currencyRef         `xml:"currency"`
		Price       zfcs_moneyType           `xml:"price"`
		Cancel      `xml:"cancel"`
	} `xml:"contract"`
	Exclude struct {
		ExcludeDate string                  `xml:"excludeDate"`
		Name        zfcs_longTextType       `xml:"name"`
		Date        string                  `xml:"date"`
		Number      zfcs_documentNumberType `xml:"number"`
		Type        string                  `xml:"type"`
	} `xml:"exclude"`
}

type zfcs_violationType struct {
	ErrCode string                  `xml:"errCode"` /* Код ошибки */
	Level   zfcs_violationLevelType `xml:"level"`   /* Тип:
	error - Ошибка
	warning - Предупреждение */
	Name        string `xml:"name"`        /* Название */
	Description string `xml:"description"` /* Описание */
}
