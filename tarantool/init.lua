local uuid = require('uuid')

box.cfg {
    listen = 3301,
    memtx_memory = 10 * 1024 ^ 3,
    wal_mode = 'none',
}

local username = 'tokens'

box.schema.user.create('tokens', { password = 'tokens', if_not_exists = true })
box.schema.user.grant('tokens', 'execute,read,write', 'universe', nil, { if_not_exists = true })

local function register_procedure(name)
    box.schema.func.create(name, { language = 'C', if_not_exists = true })
    box.schema.user.grant('tokens', 'execute', 'function', name, { if_not_exists = true })
end

register_procedure('libprocedures.get_new_token')
register_procedure('libprocedures.revoke_token')
register_procedure('libprocedures.check_token')

local tokens_space = box.schema.create_space('tokens', { if_not_exists = true })
tokens_space:create_index('pk', {
    if_not_exists = true,
    parts = { 1 },
    id = 512,
})

function get_new_token()
    local new_token = { uuid.str(), 100 }
    tokens_space:put(new_token)
    return new_token
end

function check_token(token_uuid)
    local token = tokens_space:get(token_uuid)
    if token == nil then
        return false
    end
    return true
end

function revoke_token(token_uuid)
    tokens_space:delete(token_uuid)
    print('revoke ' .. token_uuid)
    return true
end

function test_search(a, b, c)
    print(a)
    print(b)
    print(c)
end
