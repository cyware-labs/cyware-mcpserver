# 🛡️ Cyware MCP Server

<p align="center">
  <strong>A powerful Model Context Protocol (MCP) server for seamless AI integration with Cyware Products</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.24.2+-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go Version">
</p>

---

## 🚀 Overview

Cyware MCP Server is a high-performance Model Context Protocol (MCP) server built in Go, designed to provide AI agents and large language models with secure, standardized access to Cyware's cybersecurity products. This server enables seamless integration between AI systems and various Cyware applications through the standardized MCP protocol.

## ✨ Features

- **🔗 MCP Protocol Compliance**: Full implementation based on the Model Context Protocol specification
- **🎯 Multi-Application Support**: Integrated access to Cyware Intel Exchange (CTIX) and Cyware Orchestrate (CO)
- **🔒 Secure AI Integration**: Robust authentication and authorization using `config.yaml` file
- **🛠️ Tool Definitions**: Structured tools for AI agents to interact with Cyware services
- **⚙️ Configurable**: Easy configuration via YAML files
- **🚀 High Performance**: Built with Go for optimal speed and reliability

## 📁 Directory Structure

```
cyware-mcpserver/
├── 📁 applications/
│   ├── 📁 ctix/                # Cyware Intel Exchange (CTIX) MCP resources and tools
│   ├── 📁 co/                  # Cyware Orchestrate (CO) MCP resources and tools
│   └── 📁 general/             # General MCP capabilities
├── 📁 cmd/
│   ├── 📄 main.go              # MCP server entry point
│   └── 📄 config.yaml          # MCP server and application configuration
├── 📁 common/                  # Shared MCP utilities (client, config, response)
├── 📄 go.mod                   # Go module definition
├── 📄 go.sum                   # Go module dependencies
├── 📄 LICENSE                  # License file
└── 📄 README.md                # Project documentation
```

## 🏃 Getting Started

### 📋 Prerequisites

  Ensure you have the following installed:

- **Go 1.24.2 or higher** (To install Go, see https://go.dev/doc/install)
- **Access to Cyware applications** (CTIX and CO) 
- **MCP-compatible AI client** (for example, Claude, Cursor, or more) or language model integration 

### 📦 Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/cyware-labs/cyware-mcpserver.git
   cd cyware-mcpserver
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

### ⚙️ Configuration

In `cmd/config.yaml`, update the following details::
- Cyware application credentials
- MCP server transport settings — Choose either stdio or sse (with specified port)

### 🚀 Running the MCP Server

1. **Build the server:**
   ```bash
   cd cmd
   go build .
   ```

2. **Configure Claude Desktop:**

  - Quick Guide for setting up MCP on Claude: [modelcontextprotocol.io/quickstart/user](https://modelcontextprotocol.io/quickstart/user)
  - After building the server, configure the binary path and config path in the `claude_desktop_config.json` file of Claude Desktop:

   ```json
   {
     "mcpServers": {
       "cywaremcp": {
         "command": "path/to/your/binary/cmd",
         "args": [
           "-config_path",
           "path/to/your/config.yaml"
         ]
       }
     }
   }
   ```

3. Restart Claude Desktop to complete the setup and view the available Cyware MCP server tools.

# 🛠️ Available MCP Tools

## Cyware Intel Exchange (CTIX)

### Authentication & User Management
- `login-to-ctix` - Login to CTIX and generate authentication token
- `logged-in-user-details` - Get details of currently logged in user 

### CQL Query & Search
- `cql-ctix-grammar-rules` - Get CTIX CQL grammar rules
- `get-cql-query-search-result` - Run CQL query and return results

### Threat Data Management
- `get-threat-data-object-details` - Get Threat Data Object details
- `get-threat-data-object-relations` - Get Threat Data Object relations
- `get-available-relation-type` - Get available relation types

### Threat Data Bulk Actions
- `threat-data-list-bulk-action-add-tag` - Bulk add tags to threat data objects
- `threat-data-list-bulk-mark-indicator-allowed` - Bulk mark indicators as indicator allowed
- `threat-data-list-bulk-unmark-indicator-allowed` - Bulk remove indicators from indicator allowed list
- `threat-data-list-bulk-manual-review` - Bulk add threat data objects for manual review
- `threat-data-list-bulk-mark-false-positive` - Bulk mark indicators as false positive
- `threat-data-list-bulk-unmark-false-positive` - Bulk unmark indicators marked as false positives
- `threat-data-list-bulk-update-analyst-tlp` - Bulk update analyst TLP of threat data objects
- `threat-data-list-bulk-update-analyst-score` - Bulk update analyst scores of threat data objects
- `threat-data-list-bulk-deprecate` - Bulk deprecate indicators
- `threat-data-list-bulk-undeprecate` - Bulk undeprecate indicators
- `threat-data-list-bulk-add-watchlist` - Bulk add threat data objects to watchlist
- `threat-data-list-bulk-remove-watchlist` - Bulk remove threat data objects from watchlist
- `threat-data-list-bulk-add-relation` - Bulk add relation to threat data objects

### Tag Management
- `create-tag-in-ctix` - Create new tags in CTIX
- `get-ctix-tags-list` - Get list of available tags

### Enrichment Tools and Actions
- `get-enrichment-tools-list` - Get list of all enrichment tools
- `get-enrichment-tool-details` - Get details of an enrichment tool
- `get-enrichment-tool-action-configs` - Get action configuration details of enrichment tool
- `enrichment-tool-supported-for-threat-data-object` - Get supported enrichment tools for specific threat data types
- `enrich-threat-data-object` - Enrich threat data objects using configured tools

### Intel Creation
- `quick-add-intel-create` - Create intel in CTIX using Quick Add Intel

## Cyware Orchestrate (CO)

### Authentication & User Management
- `login-to-co` - Login to CO and generate the authentication token

### Playbooks Details & Execution

- `get-co-playbooks-list` - Get the list of playbooks created in CO
- `get-co-playbook-details` - Get details of a playbook
- `execute-playbook-in-co` - Run CO playbook

### CO Apps & Actions
- `get-co-apps-list` - Get the list of apps present in CO
- `get-co-app-details` _ Get the details of a specific app
- `get-co-actions-of-app` - Get list of actions supported by the app
- `get-co-app-action-details` - Get the details of an action
- `get-instances-of-co-app` - Get the instances configured in the app
- `execute-action-of-co-app` - Run action of an app

## 📄 License

This project is licensed under the terms specified in the LICENSE file.

---
