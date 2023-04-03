
generate:
	hz update -idl ./../../idl/http/$(servicename).thrift
	kitex -service user -module NihiStore -use NihiStore/server/shared/kitex_gen ../../idl/rpc/user.thrift

user:
	make servicename=user generate

goods:
	make servicename=goods generate

all:
	make user
	make goods