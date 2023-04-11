
generate:
	kitex -module NihiStore ./../idl/rpc/$(servicename).thrift

user:
	make servicename=user generate

goods:
	make servicename=goods generate

pay:
	make servicename=pay generate
all:
	make user
	make goods
	make pay