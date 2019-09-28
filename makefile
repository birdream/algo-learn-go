

rever-list:
	cd ./List && go build && ./List

quick-sort:
	cd ./Sort/quick-sort && go build -o s && ./s

radix-sort:
	cd ./Sort/radix-sort && go build -o s && ./s

insert-sort:
	cd ./Sort/insertion-sort && go build -o s && ./s

recursive-sort:
	cd ./Sort/recursive-sort && go build -o s && ./s

run-test-list:
	echo '\n==================' && \
echo '======Running test on List=====' && \
echo '==================\n' && \
cd ./List && go test