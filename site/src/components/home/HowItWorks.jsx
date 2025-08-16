
import React from "react";
import { Card, CardContent } from "@/components/ui/card";
import { ArrowRight, FileText, Terminal, Settings, RotateCcw } from "lucide-react";
import { motion } from "framer-motion";

const steps = [
  {
    icon: FileText,
    title: "shortcuts.yaml",
    description: "Parse YAML and validate shortcuts (ensures keys are allowed and no duplicates)"
  },
  {
    icon: Terminal,
    title: "KaraYAML CLI",
    description: "Generate Karabiner 'Complex Modifications' rules for each shortcut"
  },
  {
    icon: Settings,
    title: "karabiner.json",
    description: "Write the updated JSON to Karabiner's config file (in the default profile)"
  },
  {
    icon: RotateCcw,
    title: "Karabiner-Elements",
    description: "Triggers Karabiner-Elements to reload, instantly activating shortcuts. This opens the app if it's closed, or brings it to the front if it's already running."
  }
];

export default function HowItWorks() {
  return (
    <section id="how-it-works" className="bg-surface py-32">
      <div className="max-w-7xl mx-auto px-6">
        <div className="text-center mb-20">
          <motion.h2 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="text-4xl md:text-5xl font-bold mb-6"
          >
            <span className="text-text">How it works</span>
          </motion.h2>
          <motion.p 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: 0.1 }}
            className="text-xl text-text-muted max-w-3xl mx-auto"
          >
            From YAML to working keyboard shortcuts in four simple steps
          </motion.p>
        </div>

        <div className="grid lg:grid-cols-4 gap-8">
          {steps.map((step, index) => (
            <React.Fragment key={step.title}>
              <motion.div
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ delay: index * 0.15 }}
                className="relative"
              >
                <Card className="h-full bg-background border-border hover:shadow-xl transition-all duration-300">
                  <CardContent className="p-8 text-center">
                    <div className="w-20 h-20 rounded-full bg-gradient-to-br from-primary to-secondary flex items-center justify-center mb-6 mx-auto">
                      <step.icon className="w-10 h-10 text-white" />
                    </div>
                    <h3 className="text-xl font-bold mb-4 text-text">
                      {step.title}
                    </h3>
                    <p className="text-text-muted text-sm leading-relaxed">
                      {step.description}
                    </p>
                  </CardContent>
                </Card>

                {/* Arrow */}
                {index < steps.length - 1 && (
                  <div className="hidden lg:block absolute top-1/2 -right-4 transform -translate-y-1/2 z-10">
                    <div className="w-8 h-8 rounded-full bg-accent flex items-center justify-center">
                      <ArrowRight className="w-4 h-4 text-white" />
                    </div>
                  </div>
                )}
              </motion.div>
            </React.Fragment>
          ))}
        </div>
      </div>
    </section>
  );
}
