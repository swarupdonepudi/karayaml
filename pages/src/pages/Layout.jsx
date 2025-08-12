

import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { createPageUrl } from "@/utils";
import { Moon, Sun, Github, Book, Terminal } from "lucide-react";
import { Button } from "@/components/ui/button";

export default function Layout({ children, currentPageName }) {
  const [isDark, setIsDark] = useState(false);
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 10);
    };
    
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const toggleTheme = () => {
    setIsDark(!isDark);
  };

  const scrollToSection = (sectionId) => {
    const element = document.getElementById(sectionId);
    if (element) {
      element.scrollIntoView({ behavior: 'smooth' });
    }
  };

  return (
    <div className={isDark ? 'dark' : ''}>
      <style>{`
        :root {
          --primary: ${isDark ? '#3b82f6' : '#1e40af'};
          --primary-dark: ${isDark ? '#2563eb' : '#1e3a8a'};
          --secondary: ${isDark ? '#10b981' : '#059669'};
          --accent: ${isDark ? '#fbbf24' : '#f59e0b'};
          --background: ${isDark ? '#0f172a' : '#ffffff'};
          --surface: ${isDark ? '#1e293b' : '#f8fafc'};
          --border: ${isDark ? '#334155' : '#e2e8f0'};
          --text: ${isDark ? '#f1f5f9' : '#0f172a'};
          --text-muted: ${isDark ? '#94a3b8' : '#64748b'};
          --code-bg: ${isDark ? '#1e293b' : '#f1f5f9'};
        }
        
        body {
          background-color: var(--background);
          color: var(--text);
          transition: all 0.3s ease;
        }
        
        .code-block {
          background: var(--code-bg);
          border: 1px solid var(--border);
        }
        
        .surface {
          background: var(--surface);
          border: 1px solid var(--border);
        }
      `}</style>

      <div className="min-h-screen bg-background">
        {/* Navigation */}
        <nav className={`fixed top-0 w-full z-50 transition-all duration-300 ${
          isScrolled ? 'bg-background/80 backdrop-blur-lg border-b border-border' : 'bg-transparent'
        }`}>
          <div className="max-w-7xl mx-auto px-6 py-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <Terminal className="w-8 h-8 text-primary" />
                <span className="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
                  KaraYAML
                </span>
              </div>

              <div className="hidden md:flex items-center space-x-8">
                <button 
                  onClick={() => scrollToSection('why')}
                  className="text-text-muted hover:text-text transition-colors"
                >
                  Why
                </button>
                <button 
                  onClick={() => scrollToSection('how-it-works')}
                  className="text-text-muted hover:text-text transition-colors"
                >
                  How it works
                </button>
                <button 
                  onClick={() => scrollToSection('quickstart')}
                  className="text-text-muted hover:text-text transition-colors"
                >
                  Quickstart
                </button>
                <button 
                  onClick={() => scrollToSection('examples')}
                  className="text-text-muted hover:text-text transition-colors"
                >
                  Examples
                </button>
                <button 
                  onClick={() => scrollToSection('cli')}
                  className="text-text-muted hover:text-text transition-colors"
                >
                  CLI
                </button>
                <button 
                  onClick={() => scrollToSection('faq')}
                  className="text-text-muted hover:text-text transition-colors"
                >
                  FAQ
                </button>
              </div>

              <div className="flex items-center space-x-4">
                <Button
                  variant="ghost"
                  size="icon"
                  onClick={toggleTheme}
                  className="text-text-muted hover:text-text"
                >
                  {isDark ? <Sun className="w-5 h-5" /> : <Moon className="w-5 h-5" />}
                </Button>
                <a href="https://github.com/swarupdonepudi/karayaml" target="_blank" rel="noopener noreferrer">
                  <Button
                    variant="ghost"
                    size="sm"
                    className="text-text-muted hover:text-text"
                  >
                    <Github className="w-4 h-4 mr-2" />
                    GitHub
                  </Button>
                </a>
              </div>
            </div>
          </div>
        </nav>

        {/* Main Content */}
        <main className="pt-20">
          {children}
        </main>

        {/* Footer */}
        <footer className="border-t border-border bg-surface mt-32">
          <div className="max-w-7xl mx-auto px-6 py-12">
            <div className="flex flex-col md:flex-row justify-between items-center">
              <div className="flex items-center space-x-2 mb-4 md:mb-0">
                <Terminal className="w-6 h-6 text-primary" />
                <span className="text-xl font-bold">KaraYAML</span>
              </div>
              <div className="flex items-center space-x-6 text-sm text-text-muted">
                <a href="#" className="hover:text-text transition-colors">GitHub</a>
                <a href="#" className="hover:text-text transition-colors">Docs</a>
                <a href="#" className="hover:text-text transition-colors">CLI Reference</a>
                <a href="#" className="hover:text-text transition-colors">Examples</a>
                <a href="#" className="hover:text-text transition-colors">License (Apache-2.0)</a>
              </div>
            </div>
            <div className="text-center text-sm text-text-muted mt-8 pt-8 border-t border-border">
              © 2025 Swarup Donepudi. All rights reserved.
            </div>
          </div>
        </footer>
      </div>
    </div>
  );
}

