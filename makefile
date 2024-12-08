run:
	# 启动后端
	cd backend && go run main.go &
	# 启动前端
	cd frontend && npm run dev