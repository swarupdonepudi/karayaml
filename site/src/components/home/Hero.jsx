
import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Copy, Check, Github, Terminal, Zap } from "lucide-react";
import { motion } from "framer-motion";
import HeroVisual from "./HeroVisual";

export default function Hero() {
  const [copied, setCopied] = useState(false);

  const copyInstallCommand = async () => {
    await navigator.clipboard.writeText("brew install swarupdonepudi/tap/karayaml");
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  const scrollToExample = () => {
    document.getElementById('examples')?.scrollIntoView({ behavior: 'smooth' });
  };

  return (
    <div className="relative overflow-hidden">
      {/* Background gradient */}
      <div className="absolute inset-0 bg-gradient-to-br from-primary/5 via-transparent to-secondary/5 -z-10" />
      
      <div className="max-w-7xl mx-auto px-6 pt-20 pb-20">
        <div className="text-center max-w-4xl mx-auto">
          {/* Badges */}
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            className="flex flex-wrap justify-center gap-2 mb-8"
          >
            <Badge variant="outline" className="bg-surface">
              <Terminal className="w-3 h-3 mr-1" />
              Apache-2.0
            </Badge>
            <Badge variant="outline" className="bg-surface">
              <Zap className="w-3 h-3 mr-1" />
              YAML config
            </Badge>
            <Badge variant="outline" className="bg-surface">CLI-first</Badge>
            <Badge variant="outline" className="bg-surface">No dependencies</Badge>
            <Badge variant="outline" className="bg-surface">macOS</Badge>
          </motion.div>

          {/* Main heading */}
          <motion.h1 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.1 }}
            className="text-5xl md:text-7xl font-bold leading-tight mb-8"
          >
            Launch apps on Mac with Caps Lock shortcuts
          </motion.h1>

          {/* Subhead */}
          <motion.p 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.2 }}
            className="text-xl md:text-2xl text-text-muted leading-relaxed mb-8 max-w-3xl mx-auto"
          >
            Map Caps Lock + any key to open apps. Define shortcuts in one YAML file; KaraYAML updates Karabiner automatically
          </motion.p>

          <HeroVisual />

          {/* CTAs */}
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.3 }}
            className="flex flex-col sm:flex-row gap-4 justify-center items-center mb-12"
          >
            <div className="flex items-center bg-code-bg border border-border rounded-xl p-2 shadow-sm">
              <pre className="text-md font-mono text-text-muted pl-4 pr-6">
                <code>brew install swarupdonepudi/tap/karayaml</code>
              </pre>
              <Button
                onClick={copyInstallCommand}
                className="bg-primary hover:bg-primary-dark text-white rounded-lg flex-shrink-0"
                size="icon"
              >
                {copied ? (
                  <Check className="w-4 h-4" />
                ) : (
                  <Copy className="w-4 h-4" />
                )}
                <span className="sr-only">Copy install command</span>
              </Button>
            </div>
            
            <Button
              onClick={scrollToExample}
              variant="outline"
              className="px-8 py-3 text-lg font-medium rounded-xl border-2 hover:bg-surface transition-all duration-300"
              size="lg"
            >
              Try an Example
            </Button>
          </motion.div>

          {/* Tertiary links */}
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.4 }}
            className="flex justify-center gap-6 text-text-muted"
          >
            <a href="https://github.com/swarupdonepudi/karayaml" target="_blank" rel="noopener noreferrer" className="flex items-center gap-2 hover:text-text transition-colors">
              <Github className="w-4 h-4" />
              View on GitHub
            </a>
          </motion.div>
        </div>
      </div>
    </div>
  );
}
