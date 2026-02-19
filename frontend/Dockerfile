# ==========================================
# STAGE 1: Build the React/Vite Application
# ==========================================
# 1. Use Node.js to install dependencies and build the project
FROM node:20-alpine AS builder

# 2. Set the working directory
WORKDIR /app

# 3. Copy package.json and package-lock.json (or npm-shrinkwrap.json)
COPY package*.json ./

# 4. Install dependencies
RUN npm install

# 5. Copy the rest of the frontend source code
COPY . .

# 6. Build the static files (Vite usually outputs to a 'dist' folder)
RUN npm run build


# ==========================================
# STAGE 2: Serve the application with Nginx
# ==========================================
# 7. Use a lightweight Nginx image
FROM nginx:alpine

# 8. Copy the compiled static files from the builder stage
# IMPORTANT: If you use Create React App instead of Vite, change 'dist' to 'build'
COPY --from=builder /app/dist /usr/share/nginx/html

# 9. Expose port 80 for web traffic
EXPOSE 80

# 10. Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]