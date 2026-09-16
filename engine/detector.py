import sys
import json
import re
import os

AI_PATTERNS = {
    "verbose_comments": r"\/\*\*[\s\S]{200,}?\*\/",
    "generic_names": r"\$(data|result|output|temp|value|item|response)\b",
    "perfect_formatting": r"^\s{4,}\S",
    "obvious_comments": r"\/\/\s*(Increment|Initialize|Declare|Set|Get|Return)\s",
    "copilot_sig": r"@copilot|github\.com/features/copilot",
    "cursor_sig": r"cursor\.sh|cursor\.com",
    "claude_sig": r"claude\.ai|anthropic",
}

def detect_ai_code(file_path):
    if not os.path.exists(file_path):
        return {"error": "File not found"}

    with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
        content = f.read()

    evidence = []
    for pattern_name, regex in AI_PATTERNS.items():
        matches = re.finditer(regex, content, re.MULTILINE)
        for match in matches:
            evidence.append({
                "pattern": pattern_name,
                "line": content[:match.start()].count('\n') + 1,
                "snippet": match.group(0)[:80].replace('\n', ' ')
            })

    # Calculate a rough probability
    ai_probability = min(len(evidence) / 15.0, 1.0)
    
    # Determine risk level
    risk_level = "low"
    if ai_probability > 0.7: risk_level = "critical"
    elif ai_probability > 0.4: risk_level = "high"
    elif ai_probability > 0.2: risk_level = "medium"

    return {
        "file": file_path,
        "ai_probability": round(ai_probability, 2),
        "risk_level": risk_level,
        "evidence": evidence
    }

if __name__ == "__main__":
    if len(sys.argv) > 1:
        result = detect_ai_code(sys.argv[1])
        print(json.dumps(result, indent=2))
    else:
        print(json.dumps({"error": "No file provided"}))