
import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ArrowRight, CheckCircle, Heart } from "lucide-react";
import { motion } from "framer-motion";

const karabinerStrengths = [
  "Powerful, highly configurable key remapping engine",
  "Supports complex modifications, device-specific rules, and layers",
  "Active community with a large library of shared rules",
  "Rock-solid macOS integration trusted by thousands of users"
];

const karayamlAdds = [
  "Simple YAML config instead of manual JSON editing",
  "CLI-first workflow for adding, listing, searching, and filtering shortcuts",
  "Git-friendly configuration files for version control and sharing",
  "Automatic validation to prevent duplicate key mappings",
  "Instant reload -- KaraYAML updates Karabiner config and refreshes automatically"
];

export default function Comparison() {
  return (
    <section id="compare" className="bg-surface py-32">
      <div className="max-w-7xl mx-auto px-6">
        <div className="text-center mb-20">
          <motion.h2 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="text-4xl md:text-5xl font-bold mb-6"
          >
            <span className="text-text">Built on Karabiner-Elements</span>
          </motion.h2>
          <motion.p 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: 0.1 }}
            className="text-xl text-text-muted max-w-3xl mx-auto"
          >
            KaraYAML is a developer-friendly CLI layer on top of{" "}
            <a
              href="https://karabiner-elements.pqrs.org/"
              target="_blank"
              rel="noopener noreferrer"
              className="text-primary hover:underline"
            >
              Karabiner-Elements
            </a>
            , the incredible keyboard customization tool for macOS.
          </motion.p>
        </div>

        <div className="max-w-4xl mx-auto">
          {/* Gratitude card */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            className="bg-primary/5 border border-primary/20 rounded-2xl p-8 mb-12"
          >
            <div className="flex items-start gap-4">
              <Heart className="w-6 h-6 text-primary flex-shrink-0 mt-1" />
              <p className="text-text-muted leading-relaxed">
                KaraYAML exists because of the amazing work by the{" "}
                <a
                  href="https://github.com/pqrs-org/Karabiner-Elements"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-primary hover:underline font-medium"
                >
                  Karabiner-Elements
                </a>{" "}
                team. Karabiner-Elements is the engine that powers all keyboard remapping on macOS. KaraYAML simply provides a YAML and CLI interface to make configuring Caps Lock app-launch shortcuts faster and more developer-friendly.
              </p>
            </div>
          </motion.div>

          {/* Comparison cards */}
          <div className="relative">
            {/* Arrow overlay centered between the two cards (desktop) */}
            <div className="hidden lg:flex absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 z-10">
              <motion.div
                initial={{ opacity: 0, scale: 0.8 }}
                whileInView={{ opacity: 1, scale: 1 }}
                viewport={{ once: true }}
                transition={{ delay: 0.1 }}
                className="w-12 h-12 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center shadow-lg"
              >
                <ArrowRight className="w-6 h-6 text-white" />
              </motion.div>
            </div>

            <div className="grid lg:grid-cols-2 gap-6 items-stretch">
              <motion.div
                initial={{ opacity: 0, x: -20 }}
                whileInView={{ opacity: 1, x: 0 }}
                viewport={{ once: true }}
                className="flex"
              >
                <Card className="bg-blue-50/50 border-blue-200 dark:bg-blue-950/20 dark:border-blue-800 flex flex-col w-full">
                  <CardHeader>
                    <CardTitle className="text-xl text-center">
                      Karabiner-Elements
                    </CardTitle>
                    <p className="text-sm text-text-muted text-center">The foundation</p>
                  </CardHeader>
                  <CardContent className="flex-1">
                    <div className="space-y-3">
                      {karabinerStrengths.map((item, i) => (
                        <div key={i} className="flex items-start gap-3">
                          <CheckCircle className="w-5 h-5 text-blue-500 flex-shrink-0 mt-0.5" />
                          <span className="text-text-muted text-sm">{item}</span>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              </motion.div>

              {/* Arrow for mobile (between cards) */}
              <div className="flex lg:hidden justify-center py-2">
                <motion.div
                  initial={{ opacity: 0, scale: 0.8 }}
                  whileInView={{ opacity: 1, scale: 1 }}
                  viewport={{ once: true }}
                  transition={{ delay: 0.1 }}
                  className="w-10 h-10 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center"
                >
                  <ArrowRight className="w-5 h-5 text-white" />
                </motion.div>
              </div>

              <motion.div
                initial={{ opacity: 0, x: 20 }}
                whileInView={{ opacity: 1, x: 0 }}
                viewport={{ once: true }}
                transition={{ delay: 0.1 }}
                className="flex"
              >
                <Card className="bg-green-50/50 border-green-200 dark:bg-green-950/20 dark:border-green-800 flex flex-col w-full">
                  <CardHeader>
                    <CardTitle className="text-xl text-center">
                      KaraYAML adds
                    </CardTitle>
                    <p className="text-sm text-text-muted text-center">The developer experience</p>
                  </CardHeader>
                  <CardContent className="flex-1">
                    <div className="space-y-3">
                      {karayamlAdds.map((item, i) => (
                        <div key={i} className="flex items-start gap-3">
                          <CheckCircle className="w-5 h-5 text-green-500 flex-shrink-0 mt-0.5" />
                          <span className="text-text-muted text-sm">{item}</span>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              </motion.div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
