
update:
	hz update -idl ./../../idl/http/$(servicename).thrift

user:
	make servicename=user update

goods:
	make servicename=goods update

pay:
	make servicename=pay update
updateAll:
	make user
	make goods
	make pay