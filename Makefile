train_repo: 
	@python approach/repo/train.py

test_repo: 
	@python approach/repo/test.py


test_external: 
	@python approach/external/test.py

test:
	python3.11 test.py