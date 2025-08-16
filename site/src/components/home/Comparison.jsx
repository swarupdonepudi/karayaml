
import React from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { GitCompareArrows, CheckCircle, XCircle } from "lucide-react";
import { motion } from "framer-motion";

const comparisons = [
  {
    title: "Karabiner Elements UI",
    cons: [
      "Manual clicks & JSON editing",
      "Complex configuration UI",
      "Hard to version control",
      "Error-prone manual process"
    ],
    vsTitle: "KaraYAML CLI",
    pros: [
      "One YAML, auto-applied config",
      "Simple, developer-friendly syntax",
      "Git-ready configuration files",
      "Automated validation & deployment"
    ]
  }
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
            <span className="text-text">How we compare</span>
          </motion.h2>
          <motion.p 
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: 0.1 }}
            className="text-xl text-text-muted max-w-3xl mx-auto"
          >
            KaraYAML vs. the traditional approaches
          </motion.p>
        </div>

        <div className="space-y-16">
          {comparisons.map((comparison, index) => (
            <motion.div
              key={comparison.title}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: index * 0.1 }}
            >
              <div className="grid lg:grid-cols-3 gap-8 items-center">
                <Card className="bg-red-50/50 border-red-200 dark:bg-red-950/20 dark:border-red-800">
                  <CardHeader>
                    <CardTitle className="text-xl text-center">
                      {comparison.title}
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-3">
                      {comparison.cons.map((con, i) => (
                        <div key={i} className="flex items-start gap-3">
                          <XCircle className="w-5 h-5 text-red-500 flex-shrink-0 mt-0.5" />
                          <span className="text-text-muted text-sm">{con}</span>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>

                <div className="flex justify-center">
                  <div className="w-16 h-16 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center">
                    <GitCompareArrows className="w-8 h-8 text-white" />
                  </div>
                </div>

                <Card className="bg-green-50/50 border-green-200 dark:bg-green-950/20 dark:border-green-800">
                  <CardHeader>
                    <CardTitle className="text-xl text-center">
                      {comparison.vsTitle}
                    </CardTitle>
                  </CardHeader>
                  <CardContent>
                    <div className="space-y-3">
                      {comparison.pros.map((pro, i) => (
                        <div key={i} className="flex items-start gap-3">
                          <CheckCircle className="w-5 h-5 text-green-500 flex-shrink-0 mt-0.5" />
                          <span className="text-text-muted text-sm">{pro}</span>
                        </div>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
}
