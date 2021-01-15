import React, { FC, useEffect, useState } from 'react';
import "./Select.scss";

interface SelectData {
  name: string
  randId: string
}

interface props {
  value: string
  defaultValueLabel: string
  optionData: SelectData[]
  setState?: React.Dispatch<React.SetStateAction<string>>
  required?: true
}

const Select: FC<props> = (props) => {
  const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    if (props.setState) {
      props.setState(e.currentTarget.value)
    }
  }

  return (
    <select id="select" onChange={handleChange} required={props.required} value={props.value}>
      <option value="">{props.defaultValueLabel}</option>
      {props.optionData.map((option, i) => {
        return <option key={i} value={option.randId}>{option.name}</option>
      })}
    </select>
  )
}

// interface props {
//   type: "menu"
//   value: string
//   setState?: React.Dispatch<React.SetStateAction<string>>
// }

// const Select: FC<props> = (props) => {
//   const [optionValues, setOptionValue] = useState<Menu[]>([])
//   const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
//     if (props.setState) {
//       props.setState(e.currentTarget.value)
//     }
//   }

//   useEffect(() => {
//     const handleFindMenus = async () => {
//       try {
//         const response = await findMenus()
//         setOptionValue(response.menus)
//       } catch (error) {
//         console.log(error)
//       }
//     }
//     handleFindMenus()
//   }, [])

//   return (
//     <select id="select" onChange={handleChange} required={true} value={props.value} >
//       <option value="">カテゴリー</option>
//       {optionValues.map((option, i) => {
//         return <option key={i} value={option.randId}>{option.name}</option>
//       })}
//     </select>
//   )
// }

export default Select;