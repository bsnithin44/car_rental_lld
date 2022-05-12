package stations

// type StationRepo struct {
// 	Stations []Station
// }

// type IStationRepo interface {
// 	AddStation(station Station)
// 	UpdateStation(station Station)
// 	GetStation(name string) Station
// }

// var StationsDB *StationRepo

// func GetStationDb() *StationRepo {

// 	if StationsDB == nil {
// 		StationsDB := StationRepo{}
// 		return &StationsDB

// 	}
// 	return StationsDB
// }

// func (sR StationRepo) AddStation(station Station) {
// 	sR.Stations = append(sR.Stations, station)
// }

// func (sR StationRepo) GetStation(name string) *Station {
// 	for _, st := range sR.Stations {
// 		if name == st.Name {
// 			return &st
// 		}
// 	}
// 	return nil
// }

// func (sR StationRepo) UpdateStation(station Station) {

// 	for _, st := range sR.Stations {
// 		if station.Name == st.Name {
// 			st.AvailableCars = station.AvailableCars

// 		}
// 	}
// }
