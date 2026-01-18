# Troubleshooting Guide: Next.js 16 Turbopack vs. Webpack

## Proposed Titles
1. **Next.js 16: Taming the Turbopack vs. Webpack Conflict**
2. **WASM in Next.js 16: Fixing the Persistent Turbopack Build Error**
3. **The Migration Leap: Moving from Webpack to Turbopack in Next.js 16**

---

## The Problem
As of **Next.js 16**, Turbopack is enabled by default for local development. While this significantly speeds up build times, it creates a conflict if your `next.config.js` or `next.config.ts` contains a custom `webpack` configuration but lacks a `turbopack` equivalent.

If you see this error:
> ⨯ ERROR: This build is using Turbopack, with a `webpack` config and no `turbopack` config.

It means Next.js detected you are trying to customize the build, but you haven't told Turbopack how to handle those changes.

---

## Why This Happens (The "Why")
Next.js is transitioning from Webpack (the old standard) to Turbopack (the new Rust-based standard). In version 16, they've made Turbopack the default. If you have code like this in your config:

```typescript
// next.config.ts
webpack: (config) => {
  config.experiments = { asyncWebAssembly: true };
  return config;
}
```

Next.js pauses because it doesn't know if your application will break without that specific Webpack logic being replicated in Turbopack.

---

## The Solutions

### 1. The "Quick Fix" (Silence the Error)
If your app works fine without the Webpack customization (e.g., it was added by a plugin you no longer need), you can silence the error by adding an empty `turbopack` object:

```typescript
// next.config.ts
const nextConfig = {
  // ... your other config
  experimental: {
    turbopack: {}, // Silences the warning
  },
};
```

### 2. The "Rollback" (Force Webpack)
If you aren't ready to move to Turbopack yet, you can explicitly tell Next.js to use Webpack by running your dev server with a flag:

```bash
next dev --webpack
```
*Or update your `package.json`:*
```json
"scripts": {
  "dev": "next dev --webpack"
}
```

### 3. The "Correct Way" (Migrate to Turbopack)
For projects using **WebAssembly (WASM)**—common in Cardano/Mesh SDK development—you need to replicate your Webpack experiments in the new `turbopack` configuration.

**Before (Webpack only):**
```typescript
webpack: (config) => {
  config.experiments = { 
    asyncWebAssembly: true,
    layers: true 
  };
  return config;
},
```

**After (Adding Turbopack support):**
```typescript
experimental: {
  turbopack: {
    // Turbopack handles WASM differently, often requiring no extra config 
    // for standard imports, but you can define aliases or loaders here if needed.
  },
},
webpack: (config) => {
  config.experiments = { asyncWebAssembly: true, layers: true };
  return config;
},
```

> [!NOTE]
> Currently, Turbopack supports many standard Webpack features out of the box. If you are using `asyncWebAssembly`, Turbopack often works without any specific configuration inside the `turbopack: {}` block, provided you've enabled it in your project.

---

## Best Practices for Next.js 16
- **Incremental Migration**: Use `--webpack` temporarily while you test your site's functionality in the background with `--turbopack`.
- **Check Plugins**: If you didn't add the `webpack` config yourself, it might be from a library. Check if that library has a "Turbopack-ready" update.
- **WASM handling**: Ensure you are using the latest version of your SDK (like Mesh SDK) which may have built-in support for the newer build tools.
