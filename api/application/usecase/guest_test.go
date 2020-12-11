package usecase_test

// func InitMockDB() {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()
// 	mock.ExpectBegin()
// 	mock.ExpectExec("INSERT INTO guests").
// 		WithArgs(1, "test", "test", "test", "test", "test", "test", "test@test.com", "09012341234", "2020-10-10 00:00:00", "2020-10-10 00:00:00").
// 		WillReturnResult(sqlmock.NewResult(1, 11))
// 	mock.ExpectCommit()
// }

// func InitGuest() usecase.Guest {
// 	tdb, _, _ := sqlmock.New()
// 	defer tdb.Close()
// 	conn := &db.Conn{tdb}
// 	guest := repository.NewGuest(conn)
// 	responseGuest := response.NewGuest()
// 	reservation := entityx.NewReservation()
// 	repositoryReservation := repository.NewReservation(conn, reservation)
// 	usecaseGuest := usecase.NewGuest(guest, responseGuest, repositoryReservation)
// 	return usecaseGuest
// }

// func TestGuestCreate(t *testing.T) {
// 	r := &requestmodel.GuestCreate{
// 		AuthID:        "test",
// 		LastName:      "test",
// 		FirstName:     "test",
// 		LastNameKana:  "test",
// 		FirstNameKana: "test",
// 		Email:         "test@test.com",
// 		PhoneNumber:   "09012341234",
// 	}
// 	ctx := context.Background()
// 	g := InitGuest()
// 	res, _ := g.Create(ctx, r)
// 	// if res.Guest.FirstName != "test" {
// 	// 	t.Errorf("TestGuestCreate error")
// 	// }
// 	fmt.Print(res)
// }
