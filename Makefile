# --- ARMOR
.PHONY: proto-gen-armor
proto-gen-armor:
		protoc --go_out=. --go-grpc_out=. armor/proto/armor.proto

.PHONY: run-armor
run-armor:
	go run armor/cmd/grpcServer/main.go


# --- WEAPON
.PHONY: proto-gen-weapon
proto-gen-weapon:
		protoc --go_out=. --go-grpc_out=. weapon/proto/weapon.proto

.PHONY: run-weapon
run-weapon:
	go run weapon/cmd/grpcServer/main.go


# --- SHIELD
.PHONY: proto-gen-shield
proto-gen-shield:
		protoc --go_out=. --go-grpc_out=. shield/proto/shield.proto

.PHONY: run-shield
run-shield:
	go run shield/cmd/grpcServer/main.go


# --- EVANS
.PHONY: evans
evans: 
	evans -r repl 