#http -v POST :8080/cruise fromlocation:=2 tolocation:=2 crewmembers[]:=16 crewmembers[]:=28 StartDate=2024-01-01T00:00:00Z EndDate=2024-02-10T00:00:00Z
http -v PUT :8080/cruise/6 fromlocation:=2 tolocation:=2 crewmembers[]:=16 crewmembers[]:=28 StartDate=2022-01-01T00:00:00Z EndDate=2022-02-10T00:00:00Z
#http -v DELETE :8080/cruise/6
#http :8080/cruise/6
