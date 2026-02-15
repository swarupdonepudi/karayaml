
import React, { useState } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { ChevronDown, HelpCircle } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

const faqs = [
  {
    question: "Do I need to know Karabiner's JSON format?",
    answer: "No. KaraYAML handles all JSON generation — you just specify keys and apps in a YAML file."
  },
  {
    question: "Which operating systems are supported?",
    answer: "KaraYAML works on macOS only (Karabiner-Elements is macOS-only). KaraYAML is not available on Windows or Linux."
  },
  {
    question: "What keys can I use for shortcuts?",
    answer: "You can use single alphanumeric keys (a–z, 0–9) or function keys (F1–F12) as the trigger keys in your YAML."
  },
  {
    question: "Can I launch things other than apps?",
    answer: "KaraYAML is primarily for opening applications (using macOS `open`). KaraYAML launches the app if not already running and brings it to the front if it is. You can also open files or URLs, but arbitrary shell commands or keystrokes are not supported in the current version."
  },
  {
    question: "Will KaraYAML override my existing Karabiner config?",
    answer: "KaraYAML creates and updates the \"Complex Modifications\" rules in your default Karabiner profile. If you have other custom rules, back them up or merge them into your YAML, as KaraYAML manages the Caps Lock hyper-layer and its shortcuts."
  },
  {
    question: "Where is the YAML config stored?",
    answer: "KaraYAML keeps your shortcuts in `~/.kara.yaml` under your home directory. The Karabiner JSON config is updated in its standard location (`~/.config/karabiner/karabiner.json`), so Karabiner-Elements picks up the changes."
  },
  {
    question: "How do I find a specific shortcut?",
    answer: "Use `karayaml search <query>` to search by app name, `karayaml filter by-key <query>` to filter by key, or `karayaml filter by-app <query>` to filter by app path. All three support fuzzy matching by default, and filter commands accept `--exact` for exact matches."
  },
  {
    question: "Is KaraYAML free and open-source?",
    answer: "Yes -- KaraYAML is released on GitHub under the Apache-2.0 license. You can use KaraYAML freely for personal or professional purposes (contributions welcome!)."
  }
];

export default function FAQ() {
  const [openIndex, setOpenIndex] = useState(null);

  const toggleFaq = (index) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  return (
    <section id="faq" className="max-w-7xl mx-auto px-6 py-32">
      <div className="text-center mb-20">
        <motion.h2 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          className="text-4xl md:text-5xl font-bold mb-6"
        >
          <span className="text-text">Frequently asked questions</span>
        </motion.h2>
        <motion.p 
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ delay: 0.1 }}
          className="text-xl text-text-muted max-w-3xl mx-auto"
        >
          Everything you need to know about KaraYAML
        </motion.p>
      </div>

      <div className="max-w-4xl mx-auto space-y-6">
        {faqs.map((faq, index) => (
          <motion.div
            key={index}
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ delay: index * 0.05 }}
          >
            <Card className="surface border-border overflow-hidden">
              <CardContent className="p-0">
                <button
                  onClick={() => toggleFaq(index)}
                  className="w-full p-6 text-left flex items-center justify-between hover:bg-border/20 transition-colors"
                >
                  <div className="flex items-center gap-4">
                    <div className="w-8 h-8 rounded-full bg-gradient-to-r from-primary to-secondary flex items-center justify-center flex-shrink-0">
                      <HelpCircle className="w-4 h-4 text-white" />
                    </div>
                    <h3 className="text-lg font-semibold text-text">
                      {faq.question}
                    </h3>
                  </div>
                  <ChevronDown 
                    className={`w-5 h-5 text-text-muted transition-transform ${
                      openIndex === index ? 'rotate-180' : ''
                    }`} 
                  />
                </button>
                
                <AnimatePresence>
                  {openIndex === index && (
                    <motion.div
                      initial={{ height: 0, opacity: 0 }}
                      animate={{ height: 'auto', opacity: 1 }}
                      exit={{ height: 0, opacity: 0 }}
                      transition={{ duration: 0.3 }}
                      className="overflow-hidden"
                    >
                      <div className="px-6 pb-6 pl-18">
                        <p className="text-text-muted leading-relaxed">
                          {faq.answer}
                        </p>
                      </div>
                    </motion.div>
                  )}
                </AnimatePresence>
              </CardContent>
            </Card>
          </motion.div>
        ))}
      </div>
    </section>
  );
}
