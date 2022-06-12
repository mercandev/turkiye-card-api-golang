package domain

func MockData() Response {

	exampleModel := Response{

		IdentityCard: IdentityCard{
			IdentityNumber: 123456,
			Name:           "Kaan",
			Surname:        "Mercan",
			DateOfBrith:    "11.01.1900",
			Gender:         "M",
			Nationality:    "TUR",
			DrivingLicence: "B",
			IbanNumber:     "TR100000000000000000001",
			PaymentMethod:  "Troy",
		},

		Cards: []Cards{
			{
				CardType: "Debit",
				CardName: "Main",
				Bank: struct{ Banks }{
					Banks{
						BankName: "Garanti Bank",
					},
				},
			},
		},

		Transactions: []Transactions{
			{
				Company:         "Netflix",
				TransactionDate: "02.12.2000",
				TransactionHour: "15:21",
				Balance:         "-24.12",
			},

			{
				Company:         "Spotify",
				TransactionDate: "02.12.2000",
				TransactionHour: "14:10",
				Balance:         "-14.28",
			},
			{
				Company:         "Transaction",
				TransactionDate: "02.12.2000",
				TransactionHour: "13:45",
				Balance:         "+1.230.17",
			},
			{
				Company:         "Istanbul Card",
				TransactionDate: "02.12.2000",
				TransactionHour: "12:25",
				Balance:         "-50",
			},
			{
				Company:         "Ankara Card",
				TransactionDate: "02.12.2000",
				TransactionHour: "11:48",
				Balance:         "-26",
			},
		},

		Services: []Services{
			{
				ServicesName:    "Istanbul Card",
				PaymentsService: "Turkey Card",
				Balance:         "24,12",
			},
			{
				ServicesName:    "Turkhis Airlines Pass Card",
				PaymentsService: "Turkey Card",
				Balance:         "432,5",
			},
			{
				ServicesName:    "Garanti Bank Card",
				PaymentsService: "Turkey Card",
				Balance:         "2349.2",
			},
		},

		EnabledService: EnabledService{
			SafeSecurePay:  true,
			DrivingLicence: true,
			Passport:       true,
			PassengerCard:  true,
			OtherServices: []OtherServices{
				{
					Name:  "Garanti Bank",
					Value: true,
				},
				{
					Name:  "Turkish Airlines PassCard",
					Value: true,
				},
			},
		},
	}

	return exampleModel
}
