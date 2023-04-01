
generate:
	kitex -module NihiStore ./../idl/rpc/$(servicename).thrift

user:
	make servicename=user generate

all:
	make user