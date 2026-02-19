/**
 * Global Constants for the Frontend Application.
 * * VITE_API_URL is injected by Vite:
 * - In local dev: it comes from .env.local
 * - In production: it comes from Render Dashboard settings
 */

// API Backend URL
export const API_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

// Helper to check if we are running in development mode
export const IS_DEV = import.meta.env.DEV;