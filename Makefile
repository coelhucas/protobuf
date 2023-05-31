compile:
	@mkdir -p generated
	@echo "Compiling protobufs...";
	@protoc --go_opt=paths=source_relative --go_out=generated definitions.proto ||:
	@echo "Finished!"