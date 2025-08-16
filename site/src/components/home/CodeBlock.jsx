import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Copy, Check } from "lucide-react";
import { motion } from "framer-motion";

export default function CodeBlock({ code, language = "bash", filename = null, className = "" }) {
  const [copied, setCopied] = useState(false);

  const copyCode = async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <motion.div 
      initial={{ opacity: 0, y: 20 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true }}
      className={`relative group ${className}`}
    >
      {filename && (
        <div className="bg-border px-4 py-2 rounded-t-lg border border-b-0">
          <span className="text-sm text-text-muted font-mono">{filename}</span>
        </div>
      )}
      <div className={`code-block rounded-lg ${filename ? 'rounded-t-none' : ''} p-6 relative overflow-x-auto`}>
        <Button
          variant="ghost"
          size="sm"
          onClick={copyCode}
          className="absolute top-4 right-4 opacity-0 group-hover:opacity-100 transition-opacity"
        >
          {copied ? (
            <Check className="w-4 h-4 text-secondary" />
          ) : (
            <Copy className="w-4 h-4" />
          )}
        </Button>
        <pre className="text-sm font-mono text-text overflow-x-auto">
          <code>{code}</code>
        </pre>
      </div>
    </motion.div>
  );
}