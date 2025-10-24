package dealership

import (
	"api-servers/internal/models/mysql"
	"context"
	"fmt"
	"time"
)

func (s *service) GenerateSalesReport(ctx context.Context, period ReportPeriod) (*SalesReport, error) {
	allSales, err := s.sales_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get sales: %w", err)
	}

	var periodSales []mysql.Sale
	totalRevenue := float64(0)

	for _, sale := range allSales {
		if sale.Sale_Date.After(period.StartDate) && sale.Sale_Date.Before(period.EndDate) {
			periodSales = append(periodSales, sale)
			totalRevenue += sale.Sale_Price
		}
	}

	totalSales := len(periodSales)
	averageRevenue := float64(0)
	if totalSales > 0 {
		averageRevenue = totalRevenue / float64(totalSales)
	}

	vehicleSalesMap := make(map[string]*VehicleSalesData)
	for _, sale := range periodSales {
		vehicle, err := s.vehicle_repo.GetByID(sale.Vehicle_ID)
		if err == nil {
			key := fmt.Sprintf("%s-%s-%d", vehicle.Make, vehicle.Model, vehicle.Year)
			if existing, ok := vehicleSalesMap[key]; ok {
				existing.UnitsSold++
				existing.TotalRevenue += sale.Sale_Price
			} else {
				vehicleSalesMap[key] = &VehicleSalesData{
					Vehicle:      vehicle,
					UnitsSold:    1,
					TotalRevenue: sale.Sale_Price,
				}
			}
		}
	}

	var topVehicles []VehicleSalesData
	for _, data := range vehicleSalesMap {
		topVehicles = append(topVehicles, *data)
	}

	salesByStatus := make(map[string]int)
	for _, sale := range periodSales {
		salesByStatus[string(sale.Status)]++
	}

	return &SalesReport{
		Period:         period,
		TotalSales:     totalSales,
		TotalRevenue:   totalRevenue,
		AverageRevenue: averageRevenue,
		TopVehicles:    topVehicles,
		SalesByStatus:  salesByStatus,
	}, nil
}

func (s *service) GetTopPerformers(ctx context.Context, period ReportPeriod) (*PerformanceReport, error) {
	allSales, err := s.sales_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get sales: %w", err)
	}

	allSalespeople, err := s.salesperson_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get salespeople: %w", err)
	}

	if len(allSalespeople) == 0 {
		return nil, fmt.Errorf("no salespeople found")
	}

	salesByPerson := make(map[string]*SalespersonPerformance)

	for _, salesperson := range allSalespeople {
		salesByPerson[salesperson.ID] = &SalespersonPerformance{
			Salesperson:  salesperson,
			TotalSales:   0,
			TotalRevenue: 0,
			Commission:   0,
		}
	}

	for _, sale := range allSales {
		if sale.Sale_Date.After(period.StartDate) && sale.Sale_Date.Before(period.EndDate) {
			if perf, ok := salesByPerson[sale.Salesperson_ID]; ok {
				perf.TotalSales++
				perf.TotalRevenue += sale.Sale_Price
				perf.Commission += sale.Sale_Price * 0.02
			}
		}
	}

	var performanceData []SalespersonPerformance
	topSalesperson := allSalespeople[0]
	topRevenue := float64(0)

	for _, perf := range salesByPerson {
		performanceData = append(performanceData, *perf)
		if perf.TotalRevenue > topRevenue {
			topRevenue = perf.TotalRevenue
			topSalesperson = perf.Salesperson
		}
	}

	return &PerformanceReport{
		Period:          period,
		TopSalesperson:  topSalesperson,
		SalesPersonData: performanceData,
	}, nil
}

func (s *service) GetInventoryReport(ctx context.Context) (*InventoryReport, error) {
	vehicles, err := s.vehicle_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get vehicles: %w", err)
	}

	totalVehicles := len(vehicles)
	valueByMake := make(map[string]float64)
	vehiclesByStatus := make(map[string]int)
	totalAge := 0
	currentYear := time.Now().Year()

	var topValueVehicles []mysql.Vehicle

	for _, vehicle := range vehicles {
		valueByMake[vehicle.Make] += vehicle.Price
		vehiclesByStatus[string(vehicle.Status)]++
		totalAge += currentYear - vehicle.Year

		if vehicle.Price > 30000 {
			topValueVehicles = append(topValueVehicles, vehicle)
		}
	}

	averageAge := 0
	if totalVehicles > 0 {
		averageAge = totalAge / totalVehicles
	}

	return &InventoryReport{
		TotalVehicles:    totalVehicles,
		ValueByMake:      valueByMake,
		VehiclesByStatus: vehiclesByStatus,
		AverageAge:       averageAge,
		TopValueVehicles: topValueVehicles,
	}, nil
}
