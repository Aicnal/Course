# Course

## 功能需求

1. **用户认证**：
  - 注册/登录（使用教育邮箱进行认证）
  - 用户管理（学生/管理员）

2. **课程浏览与选课**：
  - 显示课程列表（搜索/分页/推荐）
  - 课程详询页（评分/点评/教师信息/开课日期）
  - 置入课程（自动化置入课程）

3. **社区互动**：
  - 课程点评
  - 讨论区

4. **数据可视化**：
  - 可视化展示课程点评
  - 趋势生成

## 技术栈

- 前端：Vue.js 3 + Vite + Element Plus (UI 框架) + Axios (API 请求)
- 后端：Gin 框架（Go）+ GORM（数据库 ORM）+ JWT（身份验证）
- 数据库：MySQL/PostgreSQL 或 MongoDB
- 部署：Docker + Nginx/Caddy

最好是全容器化，可以随时扩缩，也方便跑路（bushi）