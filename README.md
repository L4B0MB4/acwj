# acwj

Writing a compiler in go.

I use this to get more familiar with go and to learn more about compilers :)

I'm following the journey from: https://github.com/DoctorWkt/acwj

Starting at part 12 I started going in my own direction

# Concept

### Node:

```

node.{id|Root}{

    state: { //n-times
        ...
        {id}: basictype|node|basictype[]|node[]
    }

    input:(nodeparam)=>{
        return(
            <NodeId>
                <NodeId1/>
                <nodeparam>
            </NodeId>
        )
    }
}

```

## Example:

```
    node.Calculator{
        state ={
            value
        }

        input:(val1, func, val2)=>{
            return
            <Output>
                <func>
                    <val1>
                    <val2>
                </func>
            </Output>
        }
    }

    node.Add{
        input:(val1, val2)=>{
            return val1.state.value + val2.state.value
        }
    }


    node.Root{
        input:()=>{
            return
            <Calculator>
                <CalculatorValue>
                    3
                </CalculatorValue>
                <Add/>
                <CalculatorValue>
                    4
                </CalculatorValue>
            </Calculator>
        }
    }

```
