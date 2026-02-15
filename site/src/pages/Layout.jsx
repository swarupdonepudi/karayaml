

import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { createPageUrl } from "@/utils";
import { Moon, Sun, Book, Terminal } from "lucide-react";
import { Button } from "@/components/ui/button";
import GitHubStarBadge, { GitHubIcon } from "@/components/ui/GitHubStarBadge";

export default function Layout({ children, currentPageName }) {
  // Initialize from localStorage, default to light when not set
  const [isDark, setIsDark] = useState(() => {
    if (typeof window === 'undefined') return false;
    try {
      const stored = localStorage.getItem('theme');
      return stored ? stored === 'dark' : false;
    } catch (_) {
      return false;
    }
  });
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 10);
    };
    
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  // Keep the Tailwind dark class on the root element so CSS variables apply to body as well
  useEffect(() => {
    const root = document.documentElement; // <html>
    if (isDark) {
      root.classList.add('dark');
    } else {
      root.classList.remove('dark');
    }
  }, [isDark]);

  // Persist user preference
  useEffect(() => {
    try {
      localStorage.setItem('theme', isDark ? 'dark' : 'light');
    } catch (_) {
      // ignore storage errors
    }
  }, [isDark]);

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
      <div className="min-h-screen bg-background">
        {/* Navigation */}
        <nav className={`fixed top-0 w-full z-50 transition-all duration-300 ${
          isScrolled ? 'bg-background/80 backdrop-blur-lg border-b border-border' : 'bg-background/50'
        }`}>
          <div className="max-w-7xl mx-auto px-6 py-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <Terminal className="w-8 h-8 text-primary" />
                <span className="text-2xl font-bold text-text">
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
                <GitHubStarBadge repo="swarupdonepudi/karayaml" />
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
                <a href="https://github.com/swarupdonepudi/karayaml" target="_blank" rel="noopener noreferrer" className="inline-flex items-center gap-1.5 hover:text-text transition-colors">
                  <GitHubIcon className="w-4 h-4" />
                  GitHub
                </a>
                <button onClick={() => scrollToSection('cli')} className="hover:text-text transition-colors">CLI Reference</button>
                <button onClick={() => scrollToSection('examples')} className="hover:text-text transition-colors">Examples</button>
                <a href="https://github.com/swarupdonepudi/karayaml/blob/main/LICENSE" target="_blank" rel="noopener noreferrer" className="hover:text-text transition-colors">License (Apache-2.0)</a>
              </div>
            </div>
            <div className="text-center text-sm text-text-muted mt-8 pt-8 border-t border-border">
              © 2026 Swarup Donepudi. All rights reserved.
            </div>
          </div>
        </footer>
      </div>
  );
}

