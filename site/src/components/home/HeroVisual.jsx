import React from "react";
import { motion } from "framer-motion";
import { ArrowRight, Command } from "lucide-react";

const Key = ({ children, className = "" }) => (
  <div className={`flex items-center justify-center w-16 h-16 bg-surface border border-border rounded-lg shadow-sm font-mono text-lg ${className}`}>
    {children}
  </div>
);

const AppIcon = ({ icon: Icon, imageSrc, appName, showLabelOnHover = true }) => (
  <div className="relative flex flex-col items-center group">
    <div
      className="flex items-center justify-center w-16 h-16 bg-gradient-to-br from-primary/10 to-secondary/10 border border-border rounded-xl shadow-lg"
      title={appName}
      aria-label={appName}
    >
      {imageSrc ? (
        <img src={imageSrc} alt={appName} className="w-8 h-8 rounded-md" />
      ) : (
        <Icon className="w-8 h-8 text-primary" />
      )}
    </div>
    <span
      className={`absolute top-full mt-2 text-sm text-text-muted ${
        showLabelOnHover ? "hidden group-hover:block" : ""
      }`}
    >
      {appName}
    </span>
  </div>
);

export default function HeroVisual() {
  const containerVariants = {
    hidden: { opacity: 0 },
    visible: {
      opacity: 1,
      transition: {
        staggerChildren: 0.2,
        delayChildren: 0.5
      },
    },
  };

  const itemVariants = {
    hidden: { opacity: 0, y: 20 },
    visible: { opacity: 1, y: 0 },
  };
  
  const iconText = (
    <div className="flex flex-col items-center">
      <Command className="w-6 h-6" />
      <span className="text-xs font-sans mt-1">Caps Lock</span>
    </div>
  );

  return (
    <motion.div 
      variants={containerVariants}
      initial="hidden"
      whileInView="visible"
      viewport={{ once: true }}
      className="flex flex-col md:flex-row justify-center items-center gap-8 md:gap-16 mt-16 mb-12"
    >
      {/* Example 1: Slack */}
      <motion.div variants={itemVariants} className="flex items-center gap-4">
        <Key>{iconText}</Key>
        <span className="text-2xl font-light text-text-muted">+</span>
        <Key>S</Key>
        <ArrowRight className="w-8 h-8 text-secondary mx-4" />
        <AppIcon imageSrc="/slack.png" appName="Slack" />
      </motion.div>

      {/* Example 2: Cursor */}
      <motion.div variants={itemVariants} className="flex items-center gap-4">
        <Key>{iconText}</Key>
        <span className="text-2xl font-light text-text-muted">+</span>
        <Key>C</Key>
        <ArrowRight className="w-8 h-8 text-secondary mx-4" />
        <AppIcon imageSrc="/cursor.png" appName="Cursor" />
      </motion.div>
    </motion.div>
  );
}