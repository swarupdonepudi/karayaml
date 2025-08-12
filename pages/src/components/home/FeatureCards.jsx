import React from "react";
import { Card, CardContent } from "@/components/ui/card";
import { FileText, Shield, Terminal, Zap, Package } from "lucide-react";
import { motion } from "framer-motion";

const features = [
  {
    icon: FileText,
    title: "Single YAML config",
    description: "All your shortcut mappings live in one version-controlled YAML file (easy to share and track changes).",
    gradient: "from-blue-500 to-blue-600"
  },
  {
    icon: Shield,
    title: "Validation & safe defaults",
    description: "KaraYAML prevents duplicate key mappings and auto-configures Caps Lock as the Hyper modifier for you.",
    gradient: "from-green-500 to-green-600"
  },
  {
    icon: Terminal,
    title: "CLI-first workflow",
    description: "Initialize, edit, add, or list shortcuts entirely from the command line – no clunky UI, just code.",
    gradient: "from-purple-500 to-purple-600"
  },
  {
    icon: Zap,
    title: "Instant apply",
    description: "Every change updates the Karabiner config and reloads immediately, so you can iterate on shortcuts in real time.",
    gradient: "from-orange-500 to-orange-600"
  },
  {
    icon: Package,
    title: "Lightweight & no deps",
    description: "A simple Go CLI (no runtime dependencies). Install via Homebrew and use it on any Mac with Karabiner.",
    gradient: "from-pink-500 to-pink-600"
  }
];

export default function FeatureCards() {
  return (
    <section id="why" className="max-w-7xl mx-auto px-6 py-32">
      <div className="text-center mb-20">
        <motion.h2 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-4xl md:text-5xl font-bold mb-6"
        >
          Why choose{" "}
          <span className="bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            KaraYAML
          </span>
        </motion.h2>
        <motion.p 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
          className="text-xl text-text-muted max-w-3xl mx-auto"
        >
          Built for developers who want keyboard shortcuts that are maintainable, shareable, and version-controlled
        </motion.p>
      </div>

      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        {features.map((feature, index) => (
          <motion.div
            key={feature.title}
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: index * 0.1 }}
          >
            <Card className="h-full surface hover:shadow-xl transition-all duration-300 border-border group">
              <CardContent className="p-8">
                <div className={`w-16 h-16 rounded-2xl bg-gradient-to-br ${feature.gradient} flex items-center justify-center mb-6 group-hover:scale-110 transition-transform duration-300`}>
                  <feature.icon className="w-8 h-8 text-white" />
                </div>
                <h3 className="text-2xl font-bold mb-4 text-text">
                  {feature.title}
                </h3>
                <p className="text-text-muted leading-relaxed">
                  {feature.description}
                </p>
              </CardContent>
            </Card>
          </motion.div>
        ))}
      </div>
    </section>
  );
}