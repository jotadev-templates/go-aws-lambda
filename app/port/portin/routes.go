package portin

import "crm-lambda/bootstrap"

func buildRoutes(i *bootstrap.Inject) []route {

	return []route{
		{
			"GET /companies/{id}",
			i.App.Company.ExecuteFindByID,
		},
		{
			"POST /companies",
			i.App.Company.ExecuteCreate,
		},
		{
			"PUT /companies/{id}",
			i.App.Company.ExecuteUpsertAll,
		},
		{
			"DELETE /companies/{id}",
			i.App.Company.ExecuteDelete,
		},
	}

}
