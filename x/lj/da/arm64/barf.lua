local _s = string
local sub, match, gmatch, gsub = _s.sub, _s.match, _s.gmatch, _s.gsub

local splitlvl

local function splitstmt_one(c)
    if c == "(" then
        splitlvl = ")" .. splitlvl
    elseif c == "[" then
        splitlvl = "]" .. splitlvl
    elseif c == "{" then
        splitlvl = "}" .. splitlvl
    elseif c == ")" or c == "]" or c == "}" then
        if sub(splitlvl, 1, 1) ~= c then
            werror("unbalanced (), [] or {}")
        end
        splitlvl = sub(splitlvl, 2)
    elseif splitlvl == "" then
        return " \0 "
    end
    return c
end

local function splitstmt(stmt)
    -- Convert label with trailing-colon into .label statement.
    local label = match(stmt, "^%s*(.+):%s*$")
    if label then
        return ".label", { label }
    end

    -- Split at commas and equal signs, but obey parentheses and brackets.
    splitlvl = ""
    stmt = gsub(stmt, "[,%(%)%[%]{}]", splitstmt_one)
    print(gsub(stmt, "%z", "!!!"))
    if splitlvl ~= "" then
        werror("unbalanced () or []")
    end

    -- Split off opcode.
    local op, other = match(stmt, "^%s*([^%s%z]+)%s*(.*)$")
    if not op then
        werror("bad statement syntax")
    end

    -- Split parameters.
    local params = {}
    for p in gmatch(other, "%s*(%Z+)%z?") do
        params[#params + 1] = gsub(p, "%s+$", "")
    end
    if #params > 16 then
        werror("too many parameters")
    end

    params.op = op
    return op, params
end

local function printmap(m)
    for key, value in pairs(m) do
        print(key, value)
    end
end

op, params = splitstmt("mvn x1, x2, lsl #47")
print(op)
printmap(params)
